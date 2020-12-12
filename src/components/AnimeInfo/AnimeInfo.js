import { Component, Fragment } from "react";

import './AnimeInfo.css';

export default class AnimeInfo extends Component {

  //#region Properties

  state = {
    mode: 0
  }

  //#endregion

  //#region Lifecycle

  render() {
    return (
      <Fragment>
        <div className={'modal ' + (this.props.opened ? 'open' : '')}>
          <div className="modal-dialog modal-dialog-centered" role="document">
            <div className="modal-content">
              <div className="modal-header">
                <h5 className="modal-title">
                  <div className="title">
                    {this.props.anime.name}
                    {this.props.anime.year
                      ? <span className="badge badge-secondary float-right">{this.props.anime.year}</span>
                      : ''}
                  </div>
                  <p className="alt">{this.props.anime.altNames ? this.props.anime.altNames.join(",") : ''}</p>
                </h5>
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
                <ul className="nav nav-tabs">
                  <li className="nav-item">
                    <a className={'nav-link ' + (this.state.mode === 0 ? 'active' : '')}
                      onClick={() => this.onModeToggle(0)}>Openings</a>
                  </li>
                  <li className="nav-item">
                    <a className={'nav-link ' + (this.state.mode === 1 ? 'active' : '')}
                      onClick={() => this.onModeToggle(1)}>Endings</a>
                  </li>
                </ul>

                {
                  this.state.mode === 0 ?
                    <div className="tab">
                      OPS
                    </div>
                    : ''
                }

                {
                  this.state.mode === 1 ?
                    <div className="tab">
                      EDS
                    </div>
                    : ''
                }
              </div>

              <div className="modal-footer">
                <a
                  type="button"
                  className="btn btn-primary"
                  target="_blank"
                  href={`https://myanimelist.net/anime/${this.props.anime.id}`}
                >Visit MAL Page</a>
              </div>
            </div>
          </div>
        </div>
      </Fragment>
    )
  }

  //#endregion

  //#region Events

  onModeToggle(mode) {
    this.setState({ mode })
  }

  //#endregion
}
