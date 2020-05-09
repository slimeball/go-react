import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import SignIn from './pages/SignIn/SignIn'
import SignUp from './pages/SignUp/SignUp'
import BookList from './pages/BookList/BookList'
import BookAdd from './pages/BookAdd/BookAdd'

if (sessionStorage.getItem('signin')) {
  if (window.location.pathname === '/signin') {
    window.location.href = '/booklist'
  }
} else {
  if (window.location.pathname !== '/signin') {
    window.location.href = '/signin'
  }
}
export default class App extends React.Component {
  render() {
    return (
      <div className="App">
        <Router>
          <Switch>
            <Route path="/signin" component={SignIn}></Route>
            <Route path="/signup" component={SignUp}></Route>
            <Route path="/booklist" component={BookList}></Route>
            <Route path="/bookadd" component={BookAdd}></Route>
          </Switch>
        </Router>
      </div>
    );
  }
}
