// Walkhub
// Copyright (C) 2015 Pronovix
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

import React from "react";
import Step from "components/step";
import EmbedCode from "components/embedcode";
import {noop} from "form";
import {t} from "t";
import {severities} from "util";
import {Link} from "react-router";

class Walkthrough extends React.Component {

	static defaultProps = {
		walkthrough: {},
		onPlayClick: noop,
		onEditClick: noop,
		onDeleteClick: noop,
		editable: false,
		embedded: false,
		compact: false,
		linkTo: true,
	};

	static contextTypes = {
		router: React.PropTypes.func.isRequired,
	};

	render() {
		const walkthrough = this.props.walkthrough;

		const playButton = <a onClick={this.props.onPlayClick} className="btn btn-info">{t("Play")}</a>;

		if (this.props.embedded) {
			return playButton;
		}

		let counter = 0;
		const steps = walkthrough.steps ? walkthrough.steps.map((step) => {
			return <Step key={counter++} step={step} />;
		}) : null;

		const editbuttons = [];
		if (this.props.editable) {
			if (this.props.compact) {
				const href = this.context.router.makeHref("walkthrough", {uuid: walkthrough.uuid}, {});
				editbuttons.push(<a href={href} key="edit" target="_blank" className="btn btn-default">{t("Edit")}</a>);
			} else {
				editbuttons.push(<a onClick={this.props.onEditClick} key="edit" className="btn btn-default">{t("Edit")}</a>);
				editbuttons.push(<a onClick={this.props.onDeleteClick} key="delete" className="btn btn-danger">{t("Delete")}</a>);
			}
		}

		const titleName = this.props.compact && this.props.linkTo ?
			<Link to="walkthrough" params={{uuid: walkthrough.uuid}}>{walkthrough.name}</Link> :
			walkthrough.name;

		const title = (
			<div className="row">
				<div className="col-xs-10 col-md-9">
					<h2> {titleName} </h2>
				</div>
				<div className="col-xs-2 col-md-3">
					<h2>
						{playButton}
						{editbuttons}
					</h2>
				</div>
			</div>
		);

		const severity = (
			<div className="row">
				<div className="col-xs-12">
					<p> {t("This walkthrough @severity", {
						"@severity": severities[walkthrough.severity],
					})} </p>
				</div>
			</div>
		);

		const description = (
			<div className="row">
				<div className="col-xs-12">
					<p> {walkthrough.description} </p>
				</div>
			</div>
		);

		const stepsWidget = (
			<div className="row">
				<div className="col-xs-12">
					<h3> {t("Steps")} </h3>
					{steps}
				</div>
			</div>
		);

		const embed = (
			<div className="row">
				<div className="col-xs-4">
					<h3> {t("Embed code")} </h3>
				</div>
				<div className="col-xs-8">
					<EmbedCode uuid={walkthrough.uuid} />
				</div>
			</div>
		);

		return (
			<section key={walkthrough.revision} className={`walkthrough-uuid-${walkthrough.uuid} walkthrough-revision-${walkthrough.revision}`}>
				{title}
				{this.props.compact ? null : severity}
				{this.props.compact ? null : description}
				{this.props.compact ? null : stepsWidget}
				{this.props.compact ? null : embed}
			</section>
		);
	}

}

export default Walkthrough;
