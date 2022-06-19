import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './App.css';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
            <Routes>
                <Route path='/' element={Main} />
                <Route path='/login' element={Login} />
            </Routes>
        </BrowserRouter>
    </div>
  );
}

export default App;
