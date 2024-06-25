package interesting_file_names

import (
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

// TODO Unit test

func loadDefaultPatterns() []Pattern {
	patterns := make([]Pattern, 0)

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("\\.key$"),
		fileType:     "File with extension .key containing maybe sensitive information",
		priority:     interfaces.FindingPriorityHigh,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("\\.secret$"),
		fileType:     "File with extension .key containing maybe sensitive information",
		priority:     interfaces.FindingPriorityHigh,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("\\.private$"),
		fileType:     "File with extension .private containing maybe sensitive information",
		priority:     interfaces.FindingPriorityHigh,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("\\.env$"),
		fileType:     "File with extension .env containing maybe sensitive information",
		priority:     interfaces.FindingPriorityHigh,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("\\.htpasswd$"),
		fileType:     "File with extension .htpasswd containing maybe sensitive information",
		priority:     interfaces.FindingPriorityHigh,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("\\.conf$"),
		fileType:     "File with extension .conf containing maybe sensitive information",
		priority:     interfaces.FindingPriorityLow,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("\\.config$"),
		fileType:     "File with extension .config containing maybe sensitive information",
		priority:     interfaces.FindingPriorityLow,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("\\.htaaccess$"),
		fileType:     "File with extension .htaaccess$ containing maybe sensitive information",
		priority:     interfaces.FindingPriorityMedium,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("password"),
		fileType:     "Filename contains string 'password' indicating that is contains passwords",
		priority:     interfaces.FindingPriorityMedium,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("passwd"),
		fileType:     "Filename contains string 'passwd' indicating that is contains passwords",
		priority:     interfaces.FindingPriorityMedium,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("secret"),
		fileType:     "Filename contains string 'secrets' indicating that is contains sensitive information",
		priority:     interfaces.FindingPriorityMedium,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("(\\|/)shadow"),
		fileType:     "Filename is named 'shadow' indicating that it contains password hashes of a Linux system",
		priority:     interfaces.FindingPriorityHigh,
	})

	patterns = append(patterns, &pattern{
		regexPattern: regexp.MustCompile("Config(\\|/)SAM"),
		fileType:     "Filename is named 'SAM' indicating that it contains password hashes of a Windows systems",
		priority:     interfaces.FindingPriorityHigh,
	})

	return patterns
}
