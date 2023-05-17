import React from 'react'
import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom";
import List from './contacts/List';
import ViewContact from './contacts/ViewContact'; 


export default function App() {

  return (
    <Router>
      <Routes>
        <Route index element={<List />} />
        <Route path="/profile/:id" element={<ViewContact />} />
      </Routes>
    </Router>
  );
}
