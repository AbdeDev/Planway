import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import SalonDetail from './pages/salonDetail';
import SalonList from './pages/salonList';


function App() {
  console.log('App rendered');

  return (
    <Router>
      <div>
        <Routes>
          <Route path="/" element={<SalonList />} />
          <Route path="/salons/:id" element={<SalonDetail />} />
        </Routes>
      </div>
    </Router>
  );
}


export default App;
