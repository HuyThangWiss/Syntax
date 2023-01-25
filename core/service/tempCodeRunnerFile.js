const handleLogin = async (email, password) => {
    try {
        // Use the BLL layer to handle the login logic
        const isLoggedIn = await login(email, password);

        // Based on the result, show a message to the user
        if (isLoggedIn) {
            console.log("Login successful!");
        } else {
            console.log("Invalid email or password. Please try again.");
        }
    } catch (err) {
        throw err;
    }
}
