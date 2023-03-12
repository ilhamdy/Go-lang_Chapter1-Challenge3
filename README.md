# Chapter 1 Go-lang Basic (GLNG)
Go Programming Introduction

**Challenge 3**

## Biodata
> ID Regis = 1955617840-1126
> 
> Nama = Ilham Dwi Yanto
> 
> Kelas = Golang 2

<details><summary>Assignment 3</summary>
<p>

```ruby
package main

import (
	"fmt"
)

func main() {
	kalimat := "selamat malam"

	charCount := make(map[string]int)

	for _, char := range kalimat {
		charStr := string(char)
		fmt.Println(charStr)
		charCount[charStr]++
	}

	fmt.Println(charCount)
}

```
![](<Output-Challenge 3.png?raw=true>)
</p>
</details>

