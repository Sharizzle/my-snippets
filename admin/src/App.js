// import { createMuiTheme } from "@material-ui/core/styles";
import * as React from "react";
import { Admin, Resource } from "react-admin";
import { CodeList } from "./components/CodeList";
import { CodeEdit } from "./components/CodeEdit";
import { CodeCreate } from "./components/CodeCreate";
import jsonServerProvider from "ra-data-json-server";

const dataProvider = jsonServerProvider("http://localhost:3001");

const App = () => (
  <Admin dataProvider={dataProvider}>
    <Resource
      name="codes"
      list={CodeList}
      edit={CodeEdit}
      create={CodeCreate}
    />
  </Admin>
);
export default App;
