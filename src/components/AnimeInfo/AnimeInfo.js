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
                <h5 className="modal-title">
                  <div class="title">
                    {this.props.anime.name}
                    {this.props.anime.year
                      ? <span className="badge badge-secondary float-right">{this.props.anime.year}</span>
                      : ''}
                  </div>
                  <p class="alt">{this.props.anime.altNames ? this.props.anime.altNames.join(",") : ''}</p>
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
                ...
              </div>

              <div class="modal-footer">
                <a
                  type="button"
                  class="btn btn-primary"
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

  //#endregion
}
