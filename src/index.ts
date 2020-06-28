import axios from "axios";
import x from "./util";

x();
axios.get('https://github.com/axios/axios/issues/1221')
  .then(e => console.log({ success: e }))
  .catch(e => console.log({ error: e }))
