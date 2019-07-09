/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

const React = require('react');

class Footer extends React.Component {
  docUrl(doc, language) {
    const baseUrl = this.props.config.baseUrl;
    const docsUrl = this.props.config.docsUrl;
    const docsPart = `${docsUrl ? `${docsUrl}/` : ''}`;
    const langPart = `${language ? `${language}/` : ''}`;
    return `${baseUrl}${docsPart}${langPart}${doc}`;
  }

  pageUrl(doc, language) {
    const baseUrl = this.props.config.baseUrl;
    return baseUrl + (language ? `${language}/` : '') + doc;
  }

  render() {
    return (
      <footer className="nav-footer" id="footer">
        <section className="sitemap">
          <a href={this.props.config.baseUrl} className="nav-home">
            {this.props.config.footerIcon && (
              <img
                src={this.props.config.baseUrl + this.props.config.footerIcon}
                alt={this.props.config.title}
                width="66"
                height="58"
              />
            )}
          </a>
          <div>
            <h5>Docs</h5>
            <a href={this.docUrl('setup.html', this.props.language)}>
              Beginner Guides
            </a>
            <a href={this.docUrl('opspec.html', this.props.language)}>
              Opspec Reference
            </a>
            <a href={this.docUrl('interop-azure-pipelines.html', this.props.language)}>
              Interop
            </a>
          </div>
          <div>
            <h5>Community</h5>
            <a href={this.props.config.slackUrl}>Slack</a>
          </div>
          <div>
            <h5>More</h5>
            <a href={this.props.config.githubUrl}>GitHub</a>
            <a
              className="github-button"
              href={this.props.config.githubUrl}
              data-icon="octicon-star"
              data-count-href="/opctl/opctl/stargazers"
              data-show-count="true"
              data-count-aria-label="# stargazers on GitHub"
              aria-label="Star this project on GitHub">
              Star
            </a>
          </div>
        </section>

        <section className="copyright">{this.props.config.copyright}</section>
      </footer>
    );
  }
}

module.exports = Footer;
