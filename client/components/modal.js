import Reflux from 'reflux';
import React from 'react';

import ModalStore from 'stores/modalstore';

let Modal = React.createClass({
    render() {
        return (
            <div className="modal-overlay">
                <div className="modal-content">
                    {this.props.children}
                </div>
            </div>
        );
    }
});

export default Modal;
