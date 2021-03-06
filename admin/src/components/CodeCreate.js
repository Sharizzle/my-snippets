import React from "react";
import { Create, SimpleForm, DateInput, TextInput } from "react-admin";

export const CodeCreate = (props) => (
  <Create {...props}>
    <SimpleForm>
      <TextInput source="title" />
      <DateInput source="date" />
      <TextInput multiline source="description" />
      <TextInput source="category" />
      <TextInput multiline source="content" />
    </SimpleForm>
  </Create>
);
