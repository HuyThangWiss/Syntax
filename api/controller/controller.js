//showSearchByName.js
const searchByName = require('../../core/service/service');

require("dotenv").config();

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
  return result;
//  console.log(result);
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
const jwt = require('jsonwebtoken');

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


const handleLoginToken = async (username, password) => {
  try {
    const user = await searchByName.login(username, password);
    if(!user) throw new Error('Invalid username or password');
    // handle successful login
    const payload = {
      username: user.username,
    };
    const options = {
      expiresIn: '1d',
    };
    const secret ="No token provided";
    const token = jwt.sign(payload, secret, options);
    return token;
  } catch (err) {
    // handle login error
    throw err;
  }
}

const verifyToken = (req, res, next) => {
  const token = req.headers['x-access-token'];
  if (!token) {
    return res.status(403).send({
      auth: false,
      message: 'No token provided.'
    });
  }
  const secret =  "No token provided";
  jwt.verify(token,secret, (err, decoded) => {
    if (err) {
      return res.status(500).send({
        auth: false,
        message: 'Failed to authenticate token.'
      });
    }
    req.userId = decoded.id;
    next();
  });
};

module.exports = {showSearchByName,addBookPresentation,
  findAllPresentation,updateBookPresentation, 
  deleteBookPresentation,addUserPresentation,handleLogin,handleLoginToken,
  verifyToken
};
