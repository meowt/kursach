package main

var urls []string

func addUserToSlice(d string) {
	urls = append(urls, d)
}

func isUrlFree(d []string, x string) bool {
	for _, n := range d {
		if x == n {
			return false
		}
	}
	return true
}
