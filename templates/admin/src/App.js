import React from 'react';
import Header from './components/common/Navbar';
import {BrowserRouter as Router, Route} from 'react-router-dom';
import Home from './views/Home';

function App() {
  return (
    <div className="App">
      <Header></Header>
      <Router>
        <Route path="/" component={Home}></Route>
      </Router>
    </div>
  );
}

export default App;
