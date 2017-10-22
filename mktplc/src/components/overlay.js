import React, {Component} from 'react';
import History from './details';

class Overlay extends Component {
	constructor() {

		super();
	}

	close = (e) => {
		alert("fse");
	}

	render() {
		return (
			<div onClick={this.props.hideOverlay}>
			//overlay content
			//put the history here
			//pass in the props
			</div>
			)
	}

}

export default Overlay;