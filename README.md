[![Build](https://github.com/wahlflo/credentialer/actions/workflows/build.yml/badge.svg?event=push)](https://github.com/wahlflo/credentialer/actions/workflows/build.yml)
[![Tests](https://github.com/wahlflo/credentialer/actions/workflows/unittests.yml/badge.svg?event=push)](https://github.com/wahlflo/credentialer/actions/workflows/unittests.yml)
![Static Badge](https://img.shields.io/badge/Status-Alpha-blue)


# Credentialer

``Credentialer`` is a tool written in Golang for the task to scan 
large amount of files for sensitive information like private keys,
hardcoded credentials etc.

The output format of the tool can be set to ``json`` or ``csv`` which allows easy post-processing.
Also, the tool allows to resume paused scans. 

As it is hard to distinguish between real passwords and nonsense and placeholders 
the tool is implemented in such a way that it rather generates a false positive 
than ignoring a real password.
It is advisable to take a closer look at the tool output and filter out 
any obvious false positives if necessary


## Detectors

- Credential Assignments in Code: 
  - Examples: ``password:= "ABC129"`` or ``let mut password = String::from("Hallo");``
  - detection patterns based on file extensions, e.g. ``.go`` for Golang language
- File names which indicate that file contains sensitive information
  - Examples: ``cert.key`` or ``myFile.private``
- Alerting on password vaults and key storage files
  - KeePass KDB Database 1.x
  - KeePass KDBX Database 2.x
  - Bouncy Castle BKS V1 Database
  - Bouncy Castle BKS V2 Database
  - JCEKS Keystore
  - JBouncy Castle BCFKS Database
  - PKCS#12 File
- Alerting on patterns via Regex expressions for:
  - Hardcoded Credentials in Commands:
    - Example: ``mysql -u username -p"password" -h hostname database_name`` or ``sshpass -p 'MY PASSWORT' ssh admin@123.123.123.123``
    - Support for commands:
      - ``curl``
      - ``mysql``
      - ``sshpass``
  - Hardcoded Credentials in:
    - Authorization Headers, example: ``Authorization=basic  Asdhuoiads:asduhaiusd``
    - Credentials in URLs, example: ``rdp://username:password@example.com``
  - Password Hashes:
    - Argon2
    - BCrypt
    - MD5
    - SHA256
    - SHA526
  - Private Key Files:
    - in XML format, in PEM format and Putty format
    - detection and finding prioritization depending on encryption of private key
  - Public Keys
    - which are itself not a security issue but good information together with leaked private keys
  - Detections of hardcoded Tokens:
    - Amazon Marketing Services - Authentication Token
    - AWS - Access Key ID
    - AWS - Secret Key
    - Facebook - Access Token
    - GitHub - Personal Access Token
    - GitHub - OAuth Access Token
    - GitHub - App user-to-server Token
    - GitHub - App server-to-server Token
    - GitHub - App refresh token
    - GitLab - Personal Access Tokens
    - GitLab - CI/CD Job Token
    - Google Cloud - API Key
    - MailChimp - Access Token
    - Mailgun - Access Token
    - PyPi - Authentication Token
    - Slack - OAuth v2 Bot Access Token
    - Slack - OAuth v2 User Access Token
    - Slack - OAuth v2 Configuration Token
    - Slack - OAuth v2 Refresh Token
    - Slack - Webhook Token
    - Stripe - API Key
    - Telegram - Bot Authentication Token
    - Twilio - Access Token

### Something missing?
Do you miss a detector for something, feel free to create a GitHub issue.
In this case please also add how one could detect it (e.g. suggest to provide a regex or file header etc.).
It would be perfect if you also can add a pull-request in case you are capable to make the changes yourself.


## Installation

Install directly from GitHub:
```sh
go install -v github.com/wahlflo/credentialer@latest
```

## Usage
use the ``-help`` flag to view the available parameters:
```
> credentialer.exe -help
Usage of credentialer.bin:
  -debug
        displays debug information
  -failed-logMessage string
        logMessage files, which could not be scanned, to a file
  -format string
        type of output [text, json, csv], default is text (default "text")
  -help
        shows help information
  -i string
        path to the directory which should be scanned
  -llm string
        which large language model should be used for fine tuning, possible options: [none, ollama], default is none (default "none")
  -no-color
        don't use ANSI escape color codes in output
  -o string
        output file where findings will be stored, default is standard output
  -ollama-sessions int
        number of concurrent sessions, default is 1 (default 1)
  -ollama-url string
        URL of the ollama LLM, default value is http://127.0.0.1:11434 (default "http://127.0.0.1:11434")
  -resume string
        resume scan based on file containing already scanned files
  -s    suppress info messages for a clearer output
  -success-logMessage string
        logMessage scanned files to a file, needed when you want to resume a paused scan
```

### Output Formats
Multiple output formats are supported: ``text``, ``json`` and ``csv``. The default is ``text`` which is ideal
if there are not many findings, and you want to view the findings on the console.
Use the ``json`` or ``csv`` option to output the findings in a way you can process and filter them easily
after the scan finished.

The default output for the findings is the console but it can also be redirected to a file with the ``-o`` option.


### Example 1
Scan a directory and print out the findings on stdout
```
> credentialer.exe -i repo
[+] start 18 processes to scan for credentials
[+] start loading files which should be scanned
[+] loading files which should be scanned finished
-----------------------------------
New Finding: AWS - Secret Key
Priority: high
Value: wahlflo/credentialer/pkg/detectors/regex
Location: pkg\detectors\loading_basic_detectors.go
-----------------------------------
-----------------------------------
New Finding: Hardcoded MySql password parameter
Priority: high
Value: mysql -u username -p"password"
Location: pkg\detectors\regex\patterns\command_line_parameters\mysql_test.go
-----------------------------------
-----------------------------------
...
```


### Example 2
Scan a directory and write logs to files:

```
> credentialer.exe -i repo -format json -o findings.txt -success-log success.txt -failed-log fails.txt
[+] start loading files which should be scanned
[+] loading files which should be scanned finished
[+] logging scanned files to: success.txt
[+] logging scanned files which could not be scanned to: fails.txt
[+] start 23 processes to scan for credentials
[+] running for 5.0132912s; processed files: 80 / 80; remaining time: 0s
[+] finished scanning
[+] terminating
> cat success.txt
{"file_path":"pkg\\detectors\\regex\\patterns\\command_line_parameters\\curl_test.go","name":"Hardcoded MySql password parameter","priority":"high","value":"curl -u username:password https://example.com"}
{"file_path":"pkg\\detectors\\regex\\patterns\\command_line_parameters\\curl_test.go","name":"Hardcoded MySql password parameter","priority":"high","value":"curl -u username:password https://example.com\","}
{"file_path":"pkg\\detectors\\regex\\patterns\\password_hashes\\BCrypt.go","name":"AWS - Secret Key","priority":"high","value":"wahlflo/credentialer/pkg/detectors/regex"}
...
```

### Example 3
Resume a paused / interrupted scan:

```
> credentialer.exe -i repo -format json -o findings.txt -success-log success.txt -failed-log fails.txt --resume .\success.txt
[+] load previous scanned files from file: .\success.txt
[+] loaded 81 previous scanned files
[+] logging scanned files to: success.txt
[+] logging scanned files which could not be scanned to: fails.txt
[+] start 23 processes to scan for credentials
[+] start loading files which should be scanned
[+] loading files which should be scanned finished
[+] running for 5.0081236s; processed files: 0 / 0; remaining time: 0s
[+] finished scanning
[+] terminating
...
```

## Use Ollama to filter out False Positives
The ``Credentialer`` can use a Large Language Model (LLM) to filter out false positives.
The LLM response can be wrong, which can result in real credentials to be filtered out. 
The usage requires an additional setup of ollama on the local system (https://ollama.com/).

The easiest way is to use the docker container of ollama. 
The communication with ollama docker container works through an HTTP api, so you have to expose the API port ``11434`` to the host system:

```sh
docker run -d -p 11434:11434 --name ollama ollama/ollama
```

Then pull the desired LLM which should be used (e.g. llama3):

```sh
docker exec -it ollama ollama pull llama3
```

Now, the LLM is ready :smile: 

To advise ``Credentialer`` to use it you have to specify the flags ``-llm ollama``.

Example with the default settings:
```sh
credentialer.exe -i repo -llm ollama -ollama-sessions 1 -ollama-url "http://127.0.0.1:11434"
```

**Hint**: the LLM is much master if you can give the docker container access to your GPU. 
As it depends on your setup if and how this is possible, no explanation will be provided here. 


## Use in the Go Module in your project
If you want to use the project in your project you can add the module to your project:

```sh
go get github.com/wahlflo/credentialer@latest
```

```go
package main

import "github.com/wahlflo/credentialer/pkg"

func main(){
	...
	scanner := pkg.NewScanner(filesToScan, yourOutputFormatter)
	scanner.AddDetector(yourBrandNewDetector)
	scanner.StartScanning()
}
```











