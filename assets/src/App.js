//import logo from './logo.svg';
import './App.css';
import './pages/Login'
import Login from './pages/Login';
import Home from './pages/Home'
import Register from './pages/Register';
import Nav from './components/Nav';
import {BrowserRouter,Route} from 'react-router-dom'
import Post from './pages/Post';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Nav></Nav>

        <main className="form-signin"> 

            <Route path="/" exact>
              <Home></Home>
            </Route>

            <Route path="/login">
              <Login></Login>
            </Route>

            <Route path="/register">
              <Register></Register>
            </Route>

            <Route path="/post">
              <Post></Post>
            </Route>
          </main>

      </BrowserRouter>
    </div>
  );
}

export default App;
