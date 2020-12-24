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

    let collections = [];
    let tabs = [];

    if (this.props.anime && this.props.anime.collections) {

      // Populating the collections
      collections = this.props.anime.collections.filter(c => this.collectionHasThemes(c, this.state.mode)).map((col, index) => (
        <div key={index}>
          <h6>{col.name}</h6>
          <ul className="list-group mb-2">
            {
              col.themes.filter(theme => theme.type === this.state.mode).map((theme, idx) => (
                <li
                  className="list-group-item"
                  key={idx}>
                  <div>
                    <span className="badge badge-primary mr-1">{this.state.mode === 0 ? 'OP' : 'ED'} {theme.order}</span>
                    <span className="badge badge-secondary mr-1">V{theme.version}</span>
                    <b>{theme.name}</b>

                    {theme.isNSFW === true ? <span className="badge badge-danger ml-1 float-right">NSFW</span> : ''}
                    {theme.hasSpoilers === true ? <span className="badge badge-danger ml-1 float-right">SPOILERS</span> : ''}
                  </div>
                  <ul className="list-group mt-2">
                    {
                      theme.sources.map((source, i) => (
                        <li
                          className="list-group-item list-group-item-source"
                          key={i}
                        >
                          <a target="_blank"
                            href={source.link}>Video</a>
                          <audio controls>
                            <source
                              src={source.link}
                              type="audio/ogg"
                            ></source>
                          </audio>
                        </li>
                      ))
                    }
                  </ul>
                </li>
              ))
            }
          </ul>
        </div>
      ));

      // Populating the tabs
      if (this.props.anime.collections.filter(c => this.collectionHasThemes(c, 0)).length > 0) {

        // Defaulting to the openings tab
        // if (this.state.mode !== 0) {
        //   this.setState({ mode: 0 });
        // }

        tabs.push(
          <li className="nav-item">
            <a className={'nav-link ' + (this.state.mode === 0 ? 'active' : '')}
              onClick={() => this.onModeToggle(0)}>Openings</a>
          </li>
        )

      }
      // If no openings are found, switch to the endings tab
      //  else if (this.state.mode === 0) {
      //   this.setState({ mode: 1 });
      // }

      if (this.props.anime.collections.filter(c => this.collectionHasThemes(c, 1)).length > 0) {
        tabs.push(
          <li className="nav-item">
            <a className={'nav-link ' + (this.state.mode === 1 ? 'active' : '')}
              onClick={() => this.onModeToggle(1)}>Endings</a>
          </li>
        )
      }
    }

    return (
      <Fragment>
        <div className={'modal ' + (this.props.opened ? 'open' : '')}>
          {
            this.props.loading ?
              <div className="spinner spinner-border text-light float-right" role="status">
                <span className="visually-hidden"></span>
              </div> :
              <div className="modal-dialog modal-dialog-centered modal-dialog-scrollable" role="document">
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
                      {tabs}
                    </ul>

                    <div className="tab">
                      {collections}
                    </div>
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
          }
        </div>
      </Fragment >
    )
  }

  //#endregion

  //#region Events

  onModeToggle(mode) {
    this.setState({ mode })
  }

  //#endregion

  //#region Methods

  collectionHasThemes(collection, type) {
    return collection.themes.some(t => t.type === type);
  }

  //#endregion
}