// fetch("../entry_list/Hatena.json")
// .then(response => {
//     return response.json();
// })
// .then(jsondata => console.log(jsondata));

// const jsonData= require("../entry_list/Hatena.json");
// console.log(jsonData);

// import { readFileSync } from "fs";

// const jsonObject = JSON.parse(
//   readFileSync("../entry_list/Hatena.json", "utf8")
// );
// const result = {};

// jsonObject.forEach((obj) => {
//   console.log(obj);
// });

import HatenaJson from "@contents/entry_list/Hatena.json";
const HatenaEntryList = () => {
  interface Entry {
    title: string;
    url: string;
    lastUpdatedAt: string;
  }
  const obj: Entry[] = JSON.parse(HatenaJson);
  return <ul>{obj.map}</ul>;
};

export default HatenaEntryList;
