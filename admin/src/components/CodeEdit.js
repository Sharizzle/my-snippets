import React from "react";
import { DateInput, Edit, SimpleForm, TextInput } from "react-admin";

export const CodeEdit = (props) => (
  <Edit {...props}>
    <SimpleForm>
      <TextInput source="id" />
      <DateInput source="date" />
      <TextInput source="title" />
      <TextInput source="description" />
      <TextInput source="category" />
      <TextInput source="content" />
    </SimpleForm>
  </Edit>
);
