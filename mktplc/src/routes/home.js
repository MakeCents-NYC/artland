
  render() {
    return (
      <div>
        <Navbar inverse fixedTop collapseOnSelect>
          <Grid>
            <Navbar.Header>
              <Navbar.Brand>
                <a href="/"><img src={logo} weign="40" height="40"/>ArtLand</a>
              </Navbar.Brand>
            <Navbar.Toggle />
            </Navbar.Header>
          <Navbar.Collapse>
            <Navbar.Form pullLeft>
              <FormGroup>
                <FormControl type="text" placeholder="Search" />
              </FormGroup>
              {' '}
              <Button type="submit">Submit</Button>
            </Navbar.Form>
          </Navbar.Collapse>
          </Grid>
        </Navbar>
        <Jumbotron>
        </Jumbotron>
      </div>
    );
  }
}