package LC443

import "fmt"

func Compress(chars []byte) int {
	n := len(chars)
	if n == 0 {
        return 0
    }

	read, write := 0,0


	for read < n{
		count := 0
		char := chars[read]
		for read < n && chars[read] == char{
			read++
			count++
		}
		chars[write]=char
		write++


		if count > 1{
			for _, digit := range []byte(fmt.Sprintf("%d", count)){
				chars[write]= digit
				write++
			}
		}
	}
	return write
}
