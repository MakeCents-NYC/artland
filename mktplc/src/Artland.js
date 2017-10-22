import React, { Component } from 'react';

const urlForUsername = username =>
  `https://api.github.com/users/${username}`

class Artland extends Component {
  constructor(props) {
    super(props)
    this.state = {
      requestFailed: false
    }
  }

  componentDidMount() {
    fetch(urlForUsername(this.props.username))
      .then(response => {
        if (!response.ok) {
          throw Error("Network request failed")
        }

        return response
      })
      .then(d => d.json())
      .then(d => {
        this.setState({
          githubData: d
        })
      }, () => {
        this.setState({
          requestFailed: true
        })
      })
  }

  render() {

    if (this.state.requestFailed) return <p>Failed!</p>
    if (!this.state.githubData) return <p>Loading...</p>
    return (
      <div>
        <h2>{this.state.githubData.name}</h2>
      </div>
    )
  }
}

export default Artland;




// import React, { Component } from 'react';


// const urlForUsername = username =>
// `https://api.github.com/users/${username}`
// // const urlForDetails = details =>
// //   `http://localhost:3001/details`
//   //`http://artland.aestheticparamedic.club/details`

// class Artland extends Component {
//   constructor(props) {
//     super(props)
//     this.state = {
//       requestFailed: false
//     }
//   }

//   componentDidMount() {
//     //details is the arg pased in
//     fetch(urlForDetails(this.props.username))
//       .then(response => {
//         if (!response.ok) {
//           throw Error("Network request failed")
//         }

//         return response
//       })
//       .then(d => d.json())
//       .then(d => {
//         this.setState({
//           detailsData: d
//         })
//       }, () => {
//         this.setState({
//           requestFailed: true
//         })
//       })
//   }

//   render() {

//     if (this.state.requestFailed) return <p>Failed!</p>
//     if (!this.state.detailsData) return <p>Loading...</p>
//     return (
//       <div>
//         <h2>{this.state.detailsData.id}</h2>
//       </div>
//     )
//   }
// }

// export default GitHub;