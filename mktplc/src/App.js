import React, { Component } from 'react';
import { Grid, Navbar, Jumbotron, Button, FormGroup, FormControl} from 'react-bootstrap';
import logo from './logo.png';
import History from './components/details';

//import Artland from './Artland'

import './App.css';


class App extends Component {
  render() {
    return (
      <div>
        <Jumbotron>
        <History route="get"></History>
        </Jumbotron>
      </div>
    );
  }
}


export default App;



        // <Navbar inverse fixedTop collapseOnSelect>
        //   <Grid>
        //     <Navbar.Header>
        //       <Navbar.Brand>
        //         <a href="/"><img src={logo} weign="40" height="40"/>ArtLand</a>
        //       </Navbar.Brand>
        //     <Navbar.Toggle />
        //     </Navbar.Header>
        //   <Navbar.Collapse>
        //     <Navbar.Form pullLeft>
        //       <FormGroup>
        //         <FormControl type="text" placeholder="Search" />
        //       </FormGroup>
        //       {' '}
        //       <Button type="submit">Submit</Button>
        //     </Navbar.Form>
        //   </Navbar.Collapse>
        //   </Grid>
        // </Navbar>