'use strict';


/**
 * ユーザーの一覧を表示
 *
 * returns SwUserListResponse
 **/
exports.usersGET = function() {
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

