// mongodb.js
const MongoClient = require('mongodb').MongoClient;
const bcrypt = require('bcrypt');

const url = 'mongodb://localhost:27017/';
const dbName = 'Books';
const collectionName = 'Books';

let client;
let db;

const connect = () => {
  return new Promise((resolve, reject) => {
    MongoClient.connect(url, { useNewUrlParser: true, useUnifiedTopology: true }, (err, _client) => {
      if (err) {
        reject(err);
        return;
      }
      client = _client;
      db = client.db(dbName);
      resolve();
    });
  });
}

const getCollection = () => {
  return db.collection(collectionName);
}

async function findByName(name) {
  try {
    await connect();
    const collection = getCollection();
    const query = { namebooks: name };
    const result = await collection.find(query).toArray();
    client.close();
    return result;
  } catch (err) {
    throw err;
  }
}

const addBook = async (book) => {
  try {
    await connect();
    const collection = getCollection();
    const existingBook = await collection.findOne({ title: book.title });
    if (existingBook) {
      // update existing book
      const result = await collection.updateOne({ title: book.title }, { $set: book });
      console.log(`Book "${book.title}" updated successfully.`);
    } else {
      // insert new book
      const result = await collection.insertOne(book);
      console.log(`Book "${book.title}" added successfully.`);
    }
    client.close();
    return result;
  } catch (err) {
    throw err;
  }
}
const findAll = async () => {
  try {
    await connect();
    const collection = getCollection();
    const result = await collection.find({}).toArray();
    client.close();
    return result;
  } catch (err) {
    throw err;
  }
}

const updateBook = async (title, updateFields) => {
  try {
      await connect();
      const collection = getCollection();
      const result = await collection.updateOne({ title: title }, { $set: updateFields });
      client.close();
      return result;
  } catch (err) {
      throw err;
  }
}

const deleteBook = async (title) => {
  try {
      await connect();
      const collection = getCollection();
      const result = await collection.deleteOne({ title: title });
      client.close();
      return result;
  } catch (err) {
      throw err;
  }
}
const addUser = async (user) => {
  try {
      // Connect to the database
      await connect();
      const collection = getCollection();
      // Insert the user with the encrypted password
      const result = await collection.insertOne(user);
      client.close();
      return result;
  } catch (err) {
      throw err;
  }
}

const getUserByEmail = async (username) => {
  try {
    await connect();
    const collection = getCollection();
    const query = { username: username };
    const result = await collection.findOne(query);
    client.close();
    return result;
  } catch (err) {
    throw err;
  }
}

module.exports = {findByName,addBook,findAll,updateBook,deleteBook,addUser,getUserByEmail};

//module.exports = { connect, getCollection, close }
