//showSearchByName.js
const searchByName = require('../../core/service/service');

async function showSearchByName(name) {
  const result = await searchByName.searchByName(name);
  console.log(result);
}
const addBookPresentation = async (book) => {
  const result = await searchByName.addBookBLL(book);
  console.log(result);
}
const findAllPresentation = async () => {
  const result = await searchByName.findAllBLL();
  console.log(result);
}
const updateBookPresentation = async (title, updateFields) => {
  const result = await searchByName.updateBookBLL(title, updateFields);
  console.log(`Book "${title}" updated successfully.`);
  return result;
}
const deleteBookPresentation = async (title) => {
  const result = await searchByName.deleteBookBLL(title);
  console.log(`Book "${title}" deleted successfully.`);
  return result;
}

const addUserPresentation = async (username, password) => {
  const result = await searchByName.addUserBLL(username, password);
  console.log("User added successfully:", result);
  return result;
}

const handleLogin = async (email, password) => {
  try {
      // Use the BLL layer to handle the login logic
      const isLoggedIn = await searchByName.login(email, password);

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


module.exports = {showSearchByName,addBookPresentation,findAllPresentation,updateBookPresentation, deleteBookPresentation,addUserPresentation,handleLogin};
