import React, {Component} from 'react'
import { Timeline, TimelineEvent } from 'react-event-timeline'

import idea from '../images/idea.png';
import map from '../images/map.png';
import paint from '../images/paint.png';
import reward from '../images/reward.png';
import draft from '../images/demo_inprogress.png';
import draft2 from '../images/demo_inprogress2.png';
import svg from '../images/svg.svg';
import demo from '../images/demo.png';

const urlForUsername = username =>
  `https://api.github.com/users/${username}`

export default class History extends Component {
  constructor(props) {
    super(props)
    this.state = {
      requestFailed: false
    }
  }

  componentDidMount() {
    fetch(urlForUsername(this.props.username))
      .then(response => {
        if (!response.ok) {
          throw Error("Network request failed")
        }

        return response
      })
      .then(d => d.json())
      .then(d => {
        this.setState({
          githubData: d
        })
      }, () => {

        this.setState({
          requestFailed: true
        })
      })
  }

    render() {

    if (this.state.requestFailed) return <p>Failed!</p>
    if (!this.state.githubData) return <p>Loading...</p>
    return (
		    <Timeline>
		    		<TimelineEvent title={this.state.githubData.name}
	                   createdAt="2017-10-23 10:34 AM"
	                   icon={<img src={reward} />}

		            >
		                <img src={demo} />
		            </TimelineEvent>
		    		<TimelineEvent title="Loving this colour palette!"
		                           createdAt="2017-10-23 10:34 AM"
		                           icon={<img src={paint} />}
		            >
		                <img src={draft} />
		            </TimelineEvent>
		            <TimelineEvent
		                title="Work Started! It is water color medium."
		                createdAt="2017-10-23 10:26 AM"
		                icon={<img src={idea} />}
		            >
		             <img src={draft2} />
		            </TimelineEvent>
		    </Timeline>
		);
	}
} 

