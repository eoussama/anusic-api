import { Component, Fragment } from "react";
import Axios from 'axios'

import './AnimeInfo.css';

export default class AnimeInfo extends Component {

  //#region Properties

  state = {
    mode: 0,
    anime: {},
    targetTheme: null,
    loading: false
  }

  //#endregion

  //#region Lifecycle

  componentDidMount() {
    this.setState({ loading: true });

    Axios.get(`${this.props.endPoint}/anime/${this.props.animeId}`)
      .then(({ data }) => {
        this.setState({ loading: false });

        if (!data.hasError) {
          this.setState({ anime: data.data });
        }
      })
      .catch(() => {
        this.setState({ loading: false });
        this.props.onAnimeClosed();
      });
  }

  render() {

    let collections = [];
    let tabs = [];
    let player = null;

    if (this.state.anime && this.state.anime.collections) {

      // Populating the collections
      collections = this.state.anime.collections.filter(c => this.collectionHasThemes(c, this.state.mode)).map((col, index) => (
        <details key={index + (100 * (this.state.mode + 1))} open={index === 0}>
          <summary>
            <h6><span class="badge badge-light">{this.getThemesCount(col, this.state.mode)}</span> {col.name}</h6>
          </summary>

          <ul className="list-group mb-2">
            {
              col.themes.filter(theme => theme.type === this.state.mode).map((theme, idx) => (
                <li
                  className="list-group-item"
                  key={idx + (200 * (this.state.mode + 1))}>
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
                          className="list-group-item"
                          key={i + (300 * (this.state.mode + 1))}
                        >
                          <a href="#" onClick={() => this.playTheme(source.link)}>Play audio</a>
                          <a target="_blank"
                            class="ml-4"
                            href={source.link}>Open video</a>
                        </li>
                      ))
                    }
                  </ul>
                </li>
              ))
            }
          </ul>
        </details>
      ));

      // Populating the tabs
      const openingCount = this.state.anime.collections.reduce((acc, c) => acc + this.getThemesCount(c, 0), 0);
      const endingCount = this.state.anime.collections.reduce((acc, c) => acc + this.getThemesCount(c, 1), 0);

      if (this.state.anime.collections.filter(c => this.collectionHasThemes(c, 0)).length > 0) {
        tabs.push(
          <li className="nav-item">
            <a className={'nav-link ' + (this.state.mode === 0 ? 'active' : '')}
              onClick={() => this.onModeToggle(0)}><span class="badge badge-secondary">{openingCount}</span> Opening(s)</a>
          </li>
        )
      } else if (this.state.mode === 0 && !this.state.loading) {
        // If no openings are found, switch to the endings tab
        this.setState({ mode: 1 });
      }

      if (this.state.anime.collections.filter(c => this.collectionHasThemes(c, 1)).length > 0) {
        tabs.push(
          <li className="nav-item">
            <a className={'nav-link ' + (this.state.mode === 1 ? 'active' : '')}
              onClick={() => this.onModeToggle(1)}><span class="badge badge-secondary">{endingCount}</span> Ending(s)</a>
          </li>
        )
      }
    }

    if (this.state.targetTheme) {
      player = (
        <audio controls
          id="player">
          <source
            src={this.state.targetTheme}
            type="audio/ogg"
          ></source>
        </audio>
      )
    }

    return (
      <Fragment>
        <div className="modal">
          {
            this.state.loading ?
              <div className="spinner spinner-border text-light float-right" role="status">
                <span className="visually-hidden"></span>
              </div> :
              <div className="modal-dialog modal-dialog-centered modal-dialog-scrollable" role="document">
                <div className="modal-content">
                  <div className="modal-header">
                    <h5 className="modal-title">
                      <div className="title">
                        {this.state.anime.name}
                        {this.state.anime.year
                          ? <span className="badge badge-secondary float-right">{this.state.anime.year}</span>
                          : ''}
                      </div>
                      <p className="alt">{this.state.anime.altNames ? this.state.anime.altNames.join(",") : ''}</p>
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
                    {player}
                    <a
                      type="button"
                      className="btn btn-primary"
                      target="_blank"
                      href={`https://myanimelist.net/anime/${this.state.anime.id}`}
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

  /**
   * Toggles the tab mode
   * @param {number} mode The selected theme type (0 opening / 1 ending)
   */
  onModeToggle(mode) {
    this.setState({ mode })
  }

  //#endregion

  //#region Methods

  /**
   * Checks whether a collection has themes
   *
   * @param {object} collection The collection that contains the themes
   * @param {number} type The type of the theme (0 opening / 1 ending)
   */
  collectionHasThemes(collection, type) {
    return collection.themes.some(t => t.type === type);
  }

  /**
   * Gets the number of themes of a given type that belong to a collection
   *
   * @param {object} collection The collection that contains the themes
   * @param {number} type The type of the theme (0 opening / 1 ending)
   */
  getThemesCount(collection, type) {
    return collection.themes.filter(t => t.type === type && t.version === 1).length;
  }

  /**
   * Plays the clicked theme
   * @param {string} source The source of the theme (URL)
   */
  playTheme(source) {
    this.setState({ targetTheme: source });

    setTimeout(() => {
      document.getElementById('player').load();
      document.getElementById('player').play();
    }, 0);
  }

  //#endregion
}
