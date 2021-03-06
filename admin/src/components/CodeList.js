import React from "react";
import { List, Datagrid, TextField, DateField } from "react-admin";

export const CodeList = (props) => (
  <List {...props}>
    <Datagrid rowClick="edit">
      <TextField source="id" />
      <DateField source="date" />
      <TextField source="title" />
      <TextField source="description" />
      <TextField source="category" />
      <TextField source="content" />
    </Datagrid>
  </List>
);
