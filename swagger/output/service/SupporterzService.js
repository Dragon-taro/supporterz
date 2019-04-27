'use strict';


/**
 * ユーザーの一覧
 *
 * returns SwUserListResponse
 **/
exports.v1UsersGET = function() {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "users" : [ {
    "name" : "name",
    "id" : 0,
    "email" : "email"
  }, {
    "name" : "name",
    "id" : 0,
    "email" : "email"
  } ]
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * ユーザーの削除
 *
 * id Long ユーザーID
 * returns SwPostResponse
 **/
exports.v1UsersIdDELETE = function(id) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "status" : "status"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * ユーザーの詳細
 *
 * id Long ユーザーID
 * returns SwUserResponse
 **/
exports.v1UsersIdGET = function(id) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "name" : "name",
  "id" : 0,
  "email" : "email"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * ユーザーの編集
 *
 * id Long ユーザーID
 * body SwAddUser  (optional)
 * returns SwPostResponse
 **/
exports.v1UsersIdPUT = function(id,body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "status" : "status"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * ユーザーの追加
 *
 * body SwAddUser  (optional)
 * returns SwPostResponse
 **/
exports.v1UsersPOST = function(body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "status" : "status"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

