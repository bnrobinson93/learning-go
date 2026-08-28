package main

func prefixer(prefix string) func(string) string {
	return func(postfix string) string {
		return prefix + " " + postfix
	}
}
