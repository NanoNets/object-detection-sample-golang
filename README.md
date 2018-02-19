# Code Samples for Object Detection Project Written in Golang

Images and annotations taken from - https://github.com/bourdakos1/Custom-Object-Detection

Images consists of frames taken from a clip from Star Wars: The Force Awakens.
[![Watch the video](https://github.com/bourdakos1/Custom-Object-Detection/raw/master/screenshots/starwars_small.gif)](https://www.youtube.com/watch?v=xW2hpkoaIiM)

Annotations are present for each frame and have the same name as the image name. You can find the example to train a model in golang and node, by updating the api-key and model id in corresponding file. There is also a pre-processed json annotations folder that are ready payload for nanonets api, the script used is node/xml-to-json.js .


Step 1 Clone the Repo
git clone https://github.com/NanoNets/object-detection-sample-golang.git
cd object-detection-sample-golang

Step 2 Get your API Key
Get your API key here:
http://app.nanonets.com/user/api_key

Step3 Set Environment Variable
export NANONETS_API_KEY=YOUR_API_KEY_GOES_HERE

Note the golang sample uses the comverted json instead of the xml payload for convenience purposes, hence it has no dependencies.
