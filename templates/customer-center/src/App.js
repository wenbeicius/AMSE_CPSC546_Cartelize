import React from 'react';
import {BrowserRouter as Router, Route} from 'react-router-dom';
import { library } from '@fortawesome/fontawesome-svg-core';
import { faUser, faShoppingCart, faSearch } from '@fortawesome/free-solid-svg-icons';
import Header from './components/common/Navbar';
import Home from './views/Home';
import Login from './views/Login';
import Register from './views/Register';

library.add(faUser, faShoppingCart, faSearch)

function App() {
  return (
    <React.Fragment>
      <Header></Header>
      <Router>
        <Route path="/" exact component={Home}/>
        <Route path="/login" exact component={Login}/>
        <Route path="/register" exact component={Register}/>
      </Router>
    </React.Fragment>
  );
}

export default App;
