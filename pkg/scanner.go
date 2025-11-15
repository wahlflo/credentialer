package pkg

import (
	"bufio"
	"os"

	"github.com/wahlflo/credentialer/pkg/interfaces"
)

type Scanner struct {
	queuedFiles             chan interfaces.File
	scannedFiles            chan string
	errorFiles              chan string
	loadedFiles             chan interfaces.LoadedFile
	output                  interfaces.OutputFormatter
	detectors               []interfaces.Detector
	numberOfFilesScanned    safeCounter
	numberOfScannersRunning safeCounter
	llm                     interfaces.LlmConnector
}

func NewScanner(queuedFiles chan interfaces.File, output interfaces.OutputFormatter, llm interfaces.LlmConnector) *Scanner {
	return &Scanner{
		queuedFiles:  queuedFiles,
		scannedFiles: nil,
		errorFiles:   nil,
		loadedFiles:  make(chan interfaces.LoadedFile, 20),
		detectors:    make([]interfaces.Detector, 0),
		output:       output,
		llm:          llm,
	}
}

func (scanner *Scanner) SetOptionLogScannedFiles() {
	scanner.scannedFiles = make(chan string, 100)
}

func (scanner *Scanner) GetScannedFiles() chan string {
	return scanner.scannedFiles
}

func (scanner *Scanner) SetOptionLogFilesWhichCouldNotBeScanned() {
	scanner.errorFiles = make(chan string, 100)
}

func (scanner *Scanner) GetFilesWhichCouldNotBeScanned() chan string {
	return scanner.errorFiles
}

// AddDetector adds a new detector to the scanner
// so that when the scan is started the added detectors are used to alert on findings
func (scanner *Scanner) AddDetector(detector interfaces.Detector) {
	scanner.detectors = append(scanner.detectors, detector)
}

// StartScanning starts the scanning routine.
func (scanner *Scanner) StartScanning(numberOfParallelScanners int) {
	scanner.injectLargeLanguageModelIntoScanner()
	scanner.numberOfFilesScanned.Set(0)
	go scanner.preloadFilesLoop()
	scanner.startScanLoops(numberOfParallelScanners)
}

func (scanner *Scanner) injectLargeLanguageModelIntoScanner() {
	for _, x := range scanner.detectors {
		x.Inject(scanner.llm)
	}
}

func (scanner *Scanner) preloadFilesLoop() {
	for {
		fileInfo, isOpen := <-scanner.queuedFiles
		if !isOpen {
			close(scanner.loadedFiles)
			return
		}

		fileContent, err := readBinaryFile(fileInfo.GetFilepath())
		if err != nil {
			if scanner.errorFiles != nil {
				scanner.errorFiles <- fileInfo.GetFilepath()
			}
			continue
		}
		loadedFile := newLoadedFile(fileInfo, fileContent)
		scanner.loadedFiles <- loadedFile
	}
}

func (scanner *Scanner) startScanLoops(numberOfParallelScanners int) {
	counter := 0
	scanner.numberOfScannersRunning.Set(int32(numberOfParallelScanners))
	for counter < numberOfParallelScanners {
		go scanner.scanLoop()
		counter += 1
	}
}

func (scanner *Scanner) scanLoop() {
	for {
		loadedFile, channelOpen := <-scanner.loadedFiles
		if !channelOpen {
			scanner.numberOfScannersRunning.Decrement()
			return
		}

		scanner.scanFile(loadedFile)
		scanner.numberOfFilesScanned.Increment()
		if scanner.scannedFiles != nil {
			scanner.scannedFiles <- loadedFile.GetFilepath()
		}
	}
}

func (scanner *Scanner) scanFile(loadedFile interfaces.LoadedFile) {
	for _, detector := range scanner.detectors {
		_ = detector.Check(scanner.output, loadedFile)
	}
}

// IsRunning returns a boolean value if the scanner is still running
func (scanner *Scanner) IsRunning() bool {
	return scanner.numberOfScannersRunning.GetValue() > 0
}

// GetNumberOfFilesScanned returns the number of files which were already scanned
func (scanner *Scanner) GetNumberOfFilesScanned() int {
	return int(scanner.numberOfFilesScanned.GetValue())
}

func readBinaryFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size = stats.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	return bytes, err
}
