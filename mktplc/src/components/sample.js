import React, { Component } from 'react';

import Gallery from 'react-photo-gallery';
import one from '../images/1.jpg';
import two from '../images/2.jpg';
import thr from '../images/3.jpg';
import fou from '../images/4.jpg';
import fiv from '../images/5.jpg';
import six from '../images/6.jpg';




export default class Sample extends Component {
  // props is a list of photos
  constructor(props) {
    super(props);

    this.state = {
       showComponent: false,
    };

    this.popup_ques = this.popup_ques.bind(this);
    this.hide_overlay = this.hide_overlay.bind(this);
  }

  popup_ques() {
    this.setState({
      showComponent: true,
    });
  }

  hide_overlay() {
    this.setState({
      showComponent: false
    })
  }

	render() {
		return (
      <div>
			    <Gallery photos={ PHOTO_SET } onClick={this.popup_ques}/>
	    </div>

      {this.state.showComponent && <Overlay hideOverlay={this.hide_overlay} />}
   );
	}
} 

const PHOTO_SET = [
  {
    src: one,
  },
  {
    src: two
  },
  {
    src: thr
  },
  {
    src:  fou
  },
  {
    src: fiv
  },
  {
    src: six
  }
];
