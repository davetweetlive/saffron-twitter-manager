package main

import "fmt"

func main() {
	// name := "Dave Augustus"

	tpl := `
	<!DOCTYPE html>
	<head>

	</head>
	<html>
	<body>
        <div class="login-box">
			<h2>Login with twitter</h2>
			<!-- <a href="#" class="social-button" id="facebook-connect"> <span>Connect with Facebook</span></a> -->
			<!-- <a href="#" class="social-button" id="google-connect"> <span>Connect with Google</span></a> -->
			<a href="#" class="social-button" id="twitter-connect"> <span>Connect with Twitter</span></a>
			<!-- <a href="#" class="social-button" id="linkedin-connect"> <span>Connect with LinkedIn</span></a> -->
		</div>
	</body>
	</html>
	`

	fmt.Println(tpl)
}
