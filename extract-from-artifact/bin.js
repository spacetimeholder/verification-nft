#!/usr/bin/env node

const yargs = require("yargs");
const artifact = yargs.argv.artifact;
const bin = yargs.argv.bin;
const abi = yargs.argv.abi;

const fs = require("fs");
let raw = fs.readFileSync(artifact);
let json = JSON.parse(raw);

var buff = Buffer.alloc(json.bytecode.length, json.bytecode, "utf8");
fs.writeFileSync(bin, buff);

const abiJson = JSON.stringify(json.abi);
buff = Buffer.alloc(abiJson.length, abiJson, "utf8");
fs.writeFileSync(abi, buff);
