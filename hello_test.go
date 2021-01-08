/*
Writing a test is just like writing a function, with a few rules
	- It needs to be in a file with a name like xxx_test.go
	- The test function must start with the word Test
	- The test function takes one argument only t *testing.T
*/

func TestHello(t *testing.T) {

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
    })

    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
    })

	//We need to pass in t *testing.T so that we can tell the test code to fail when we need to.
	//t.Helper() is needed to tell the test suite that this method is a helper. 
	//By doing this when it fails the line number reported will be in our function call rather than inside our test helper. 
    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }
}