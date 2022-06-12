# covid-spread.viz

<img src="https://user-images.githubusercontent.com/46757971/172483410-9069e2e0-c2e9-483c-b645-b4674f951d1b.gif" width=60% height=60%>

[**Live Website**](https://covidviz.com)

**covid-spread.viz**'s goal is to deliver a fast and interactive visualizer for the historical data of COVID-19 since the outbreak to the present day. It was built in order for us to better understand the tragic impact of the coronavirus pandemic over the world population. 

It is built using React.js, GoLang, Mapbox, and gRPC. The web application and the back-end server are not fully open-sourced, but the repository is semi-public to introduce the general project structure and the tech stack. See [Open Source](#open-source) section for more details.


## Tech Stack

<img src="https://user-images.githubusercontent.com/46757971/173222371-538e6579-1e4f-47b5-824e-75437929e7c1.png" width=60% height=60%>

- Data: MongoDB Atlas, JHU CSSE Covid-19 DataSet
- Back-end: GoLang, Google Cloud Run, gRPC, gRPC-web, Protocol Buffers
- Front-end: React.js, TypeScript, Mapbox, Google App Engine, gRPC-web

## Data
The original data source is John Hopkins University CSSE's [Covid-19 DataSet](https://github.com/CSSEGISandData/COVID-19/tree/master/csse_covid_19_data).
Instead of directly parsing the CSV data from JHU CSSE, the data-fetching microservice pulls data from MongoDB Atlas that has unmodified, frequently updated copy of the JHU data. Check this [blog post](https://www.mongodb.com/developer/products/atlas/johns-hopkins-university-covid-19-data-atlas/) for more details about how MongoDB hosts a free service that provides COVID-19 dataset.

Note that recovery data for United States is absent since December 14, 2020, as well as recovery data for all countries since August 5th, 2021.
JHU CSSE stopped tracking recovery data without providing much context, but you can check out the GitHub issue [here](https://github.com/CSSEGISandData/COVID-19/issues/4465).

## Open Source
This repository is not fully open-sourced (yet). Some of the major front-end and back-end files are private, but most of the contents are public to give you a general idea about the project structure and the tech stack. I am willing to open-source the whole repository if there are a good amount of people requesting for open source. If you want to learn about the details of this repository, you can vote for open-source [here](https://covid-spread-viz.canny.io/requests/p/open-source).

## TODO
- Add population data on country hover
- Add case-population ratio as a data option
- Add active vases (confirmed-recovery-deaths) as a data option, although most of the recovery data is absent
