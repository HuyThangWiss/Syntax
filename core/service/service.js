const findByName = require('../../adapter/db');



async function searchByName(name) {
  const result = await findByName.findByName(name);
  return result;
}


const addBookBLL = async (book) => {
    const result = await findByName.addBook(book);
    return result;
}

const findAllBLL = async () => {
    const result = await findByName.findAll();
    return result;
}

const updateBookBLL = async (title, updateFields) => {
    const result = await findByName.updateBook(title, updateFields);
    return result;
}
const deleteBookBLL = async (title) => {
    const result = await findByName.deleteBook(title);
    return result;
}

const bcrypt = require('bcrypt');
const saltRounds = 10;

const encryptPasswordBLL = (password) => {
    // Use bcrypt to encrypt the password
    const encryptedPassword = bcrypt.hashSync(password, saltRounds);
    return encryptedPassword;
}


const addUserBLL = async (username, password) => {
    const encryptedPassword = encryptPasswordBLL(password);
    const result = await findByName.addUser({username: username, password: encryptedPassword});
    return result;
}
//password
const login = async (email, password) => {
    const user = await findByName.getUserByEmail(email);
    if(user){
      const match = await bcrypt.compare(password, user.password);
      if(match){
        return user;
      }else{
        throw new Error('Invalid email or password');
      }
    }else{
      throw new Error('Invalid email or password');
    }
}



module.exports = {searchByName,addBookBLL,findAllBLL,updateBookBLL,deleteBookBLL,addUserBLL,login};