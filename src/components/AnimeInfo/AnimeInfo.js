import { Component, Fragment } from "react";

import './AnimeInfo.css';

export default class AnimeInfo extends Component {

  //#region Properties

  state = {}

  //#endregion

  //#region Lifecycle

  render() {
    return (
      <Fragment>
        <div className={'modal ' + (this.props.opened ? 'open' : '')} tabindex="-1">
          <div className="modal-dialog modal-dialog-centered" role="document">
            <div className="modal-content">
              <div className="modal-header">
                <h5 className="modal-title" id="exampleModalLongTitle">Modal title</h5>
                <button
                  type="button"
                  className="close"
                  data-dismiss="modal"
                  aria-label="Close"
                  onClick={() => this.props.onAnimeClosed()}>
                  <span aria-hidden="true">&times;</span>
                </button>
              </div>
              <div className="modal-body">
                ...
              </div>
            </div>
          </div>
        </div>
      </Fragment>
    )
  }

  //#endregion

  //#region Events

  //#endregion
}
