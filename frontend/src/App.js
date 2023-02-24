import { Routes, Route } from 'react-router-dom';
import React from 'react';
import './index.css';

import Home from './pages/home';

const App = () => (
  <div className="app">
    <div className='content'>
      <Routes>
        <Route exact path='/' element={<Home />}></Route>
      </Routes>
    </div>
  </div>
);

export default App;