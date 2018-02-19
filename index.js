'use strict';

const dns = require('dns');

dns.resolveMx('ya.ru', function(e,r){
  if(e){
    console.log(e);
  }
  console.log(r);
});