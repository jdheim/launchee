/*
 * © 2025-2025 JDHeim.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import type {ReactNode} from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

type FeatureItem = {
  title: string;
  image: string;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Launch anything with a command',
    image: require('@site/static/img/feature1.jpg').default,
    description: (
      <>
          Tell Launchee what to run, and it opens instantly.
          Point it at a binary, script, or alias, and you’re done.
      </>
    ),
  },
  {
    title: 'Open the web right where you need it',
    image: require('@site/static/img/feature2.jpg').default,
    description: (
      <>
          Start your browser with a specific URL in one click.
          Jump straight into dashboards, docs, or any site without typing.
      </>
    ),
  },
  {
    title: 'Simple YAML configuration',
    image: require('@site/static/img/feature3.jpg').default,
    description: (
      <>
          A clean YAML file powers everything.
          Add a name, an icon, and a command (or URL) and it's live.
          Easy to version, easy to share.
      </>
    ),
  },
];

function Feature({title, image, description}: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <img src={image} alt={title} className={styles.featureImg} />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
