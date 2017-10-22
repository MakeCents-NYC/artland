import React, { Component } from 'react';
import Gallery from 'react-photo-gallery';
import one from '../images/1.jpg';
import two from '../images/2.jpg';
import thr from '../images/3.jpg';
import fou from '../images/4.jpg';
import fiv from '../images/5.jpg';
import six from '../images/6.jpg';


export default class Sample extends Component {
	render() {
		return (
			<Gallery photos={ PHOTO_SET } onClick=details()/>
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
