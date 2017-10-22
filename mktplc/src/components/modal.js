//derived from https://daveceddia.com/open-modal-in-react/


import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {ReactModal} from 'react-modal';

class Modal extends Component {
	constructor () {
		super();
		this.state = {
			showModal: false
		};

		this.handleOpenModal = this.handleOpenModal.bind(this);
		this.handleCloseModal = this.handleCloseModal.bind(this);
	}

	handleOpenModal () {
		this.setState({ showModal: true });
	}

	handleCloseModal () {
		this.setState({ showModal: false });
	}

	render () {
		return (
		      <div>
		        <button onClick={this.handleOpenModal}>Trigger Modal</button>
		        <ReactModal 
		           isOpen={this.state.showModal}
		           contentLabel="Minimal Modal Example"
		        >
		          <button onClick={this.handleCloseModal}>Close Modal</button>
		        </ReactModal>
		      </div>
	      );

	}

}

const props = {
	  onClose: PropTypes.func.isRequired,
      show: PropTypes.bool,
	  children: PropTypes.node
}



// class Modal extends React.Component {
//   render() {
//     // Render nothing if the "show" prop is false
//     if(!this.props.show) {
//       return null;
//     }

//     // The gray background
//     const backdropStyle = {
//       position: 'fixed',
//       top: 0,
//       bottom: 0,
//       left: 0,
//       right: 0,
//       backgroundColor: 'rgba(0,0,0,0.3)',
//       padding: 50
//     };

//     // The modal "window"
//     const modalStyle = {
//       backgroundColor: '#fff',
//       borderRadius: 5,
//       maxWidth: 500,
//       minHeight: 300,
//       margin: '0 auto',
//       padding: 30
//     };

//     return (
//       <div className="backdrop" style=>
//         <div className="modal" style=>
//           {this.props.children}

//           <div className="footer">
//             <button onClick={this.props.onClose}>
//               Close
//             </button>
//           </div>
//         </div>
//       </div>
//     );
//   }
// }

// Modal.propTypes = {
//   onClose: PropTypes.func.isRequired,
//   show: PropTypes.bool,
//   children: PropTypes.node
// };

// export default Modal;