var controller = require('../../MonGoDb/api/controller/controller')


// console.log("Find")
// controller.showSearchByName('Book1').then((result) => {
//     console.log(result);
// }).catch((error) => {
//     console.log(error);
// });

/*
const data = {title: "The Great Gatsby 002", author: "F. Scott Fitzgerald 03", year: 129,address:"New york 22"}

controller.addBookPresentation(data);
*/
// console.log("Find ByName");
// const data2 = controller.showSearchByName('Books');
// console.log(data2);
/*
console.log("All data");
controller.findAllPresentation()
*/
/*
const data = {title: "The Great Gatsby 002", author: "F. Scott Fitzgerald 03", year: 129,address:"New york 22",Point:99888}

const gg =controller.addBookPresentation(data);
console.log(gg);
*/
/*
controller.updateBookPresentation("The Great Gatsby 002",{
    "author" : "F. Scott Fitzgerald 03",
    "year" : 1292,
    "address" : "New york 223",
    "Point" : 998883
})
*/

//controller.deleteBookPresentation("The Great Gatsby 002")

//const data = {title: "The Great Gatsby 002", author: "F. Scott Fitzgerald 03", year: 129,address:"New york 22",Point:99888,UserName:"Thang",Passworld:1234}

//controller.addBookPresentation(data);


//const user =  controller.addUserPresentation("john", "mysecretpassword");



const handleLogin = async (email, password) => {
    try {
      const user = await controller.handleLogin(email, password);
      // handle successful login
    } catch (err) {
      // handle login error
    }
  }
  handleLogin("john", "mysecretpassword");




















