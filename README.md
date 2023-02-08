# gpt3-rental

Use GPT-3 model (can understand and generate natural language) to generate description for the provided rental id (fetch rental details from Outdoorsy staging environment).  
In order to make more compelling description we ask GPT-3 to provide an idea for camping activity. 
We also use the National Parks Service API to add information for nearly (campgrounds and rental are in the same state) located campgrounds.

**Note: more than 90% of the source code was generated using ChatGPT**

## Prerequisites

* OpenAI API Key - https://openai.com/api/
* National Parks Service API Key - https://www.nps.gov/subjects/developer/api-documentation.htm
* Access to Outdoorsy APIs - https://search.outdoorsy.com/rentals/322681

## Usage

* setup required env variables ***NPS_API_KEY*** and ***GPT3_API_KEY***
* run the application ```go run cmd/main.go```
* access the API http://localhost:8081/<rental_id>

## Pricing

We use OpenAI API which is paid (https://openai.com/api/pricing/). The price with the model (Davinci) we use will be around $0.0200 / 1K tokens. (You can think of tokens as pieces of words, where 1,000 tokens is about 750 words. )

## Sample Descriptions

* Go on a journey of a lifetime in our Adirondack 24rgb-M5 RV rental. This Class B RV is perfect for both your weekend getaways and longer journeys. With its 1984 Volkswagen Westfalia, you can have an enjoyable and comfortable stay everywhere you go. The RV is equipped with modern amenities, such as a inside shower, toilet, generator, audio inputs, CD player, radio, TV/DVD, dining table, kitchen sink, microwave, oven, refrigerator, stove, awning, extra storage, air conditioner, ceiling fan, heating, and leveling jacks. This RV also comes with an extra storage and awning. Take on an unforgettable adventure in a unique way and explore the natural wonders of the Otis Pike High Dune Wilderness, or the unforgettable view of the Fort Wadsworth Campground. Experience the freedom of the open road in our RV rental.
* This 1984 Volkswagen Westfalia rv rental is perfect for a memorable outdoor excursion. Loaded with amenities like an inside shower and toilet, audio inputs, awning, air conditioner, and extra storage, you'll have all the comfort you need for your backcountry adventure. Explore the Otis Pike High Dune Wilderness with a camping permit and pitch your tent under the stars for up to 3 nights. Wake up to the sound of songbirds for a peaceful camping experience. For a more post-modern adventure, Fort Wadsworth Campground offers modern campsites with conveniences like a kitchen sink, refrigerator, and oven. During your stay, why not take a scenic hike around the area or take a dip in the nearby ocean?
* This 1984 Volkswagen Westfalia rental can turn any camping trip into an adventure! With a fully equipped kitchen, inside shower, and all the comforts of home, you'll have no trouble enjoying all the great camping activities available while off-roading into the Otis Pike High Dune Wilderness. And don't forget to set up the awning and leveling jacks to stay comfortable while stargazing, beachcombing, or exploring all that Fire Island has to offer! For an added surprise, crank up the tunes with the radio or CD player, or watch a movie on the TV/DVD – the perfect way to end a long day of wilderness camping!
* Rent this 1984 Volkswagen Westfalia and take advantage of the flexible camping experience it offers! With an inside shower, toilet, generator, audio inputs, CD player, radio, TV / DVD, dining table, kitchen sink, microwave, oven, refrigerator, stove, awning, extra storage, leveling jacks, air conditioner, ceiling fan, and heater, you'll have everything you need for an unforgettable camping adventure. Enjoy a unique camping experience in the beautiful Adirondack region, and don't forget to bring your beach gear for an unforgettable backcountry camping experience in the Otis Pike High Dune Wilderness. Enjoy the night sky, splash in the surf, and sing along with the birds. For added fun, why not bring a kayak and explore the Fort Wadsworth Campground on your next camping trip!
* Rent this comfortable Class B 1984 Volkswagen Westfalia that's perfect for your adventure. This vintage RV has all the amenities and features to make you feel at home during your travels. Experience the convenience of having a shower, toilet and kitchen facilities. With an audio input, CD player, radio and TV/DVD, you can stay entertained for hours. An air conditioner, ceiling fan and heater will keep you comfortable, and the awning and extra storage will give you more room. Plus, there is a generator on board and levelers, so you're all set! You can take this RV anywhere, with no outside shower, no Wi-Fi, and no satellite. Yet, it's conveniently located near campgrounds in Brooklyn, NY. With a flexible rental, you'll have no problem renting at a great price!
* This 1984 Volkswagen Westfalia is the perfect RV rental for a family excursion! With its flexible floor plan, this RV is perfect for those with a sense of adventure. It includes plenty of features like an inside shower, a toilet, a generator, CD player, radio, TV/DVD, dining table, kitchen sink, microwave, oven, refrigerator, stove, awning, extra storage, and leveling jacks. Enjoy all of the comforts of home with its air conditioner, ceiling fan, and heater. It can sleep up to 5 comfortably, and it is located in Brooklyn, US, 40.693 -73.933, NY 11221. With plenty of nearby campgrounds, there's something for everyone! Opt for Backcountry Camping in the Otis Pike High Dune Wilderness where you can explore the great outdoors under the stars near the sound of the surf. Or choose Fort Wadsworth Campsite where you can go camping with your family or a group of friends. Sandy Hook Camp Ground offers 20 tent campsites for seasonal camping, or Watch Hill Family Campground for a one-of-a-kind camping experience. Regardless of your destination, this RV rental is an ideal choice for fun and adventure.
* This 1984 Volkswagen Westfalia is the perfect rental to experience Brooklyn, NY and the surrounding area. This RV has the home comforts you need to take your adventure on the road with an inside shower, toilet, generator, audio inputs, CD player, radio, TV/DVD, dining table, kitchen sink, microwave, oven, refrigerator, stove, awning, extra storage, leveling jacks, air conditioner, ceiling fan, and heater. Even though there are no external amenities as a backup camera, bike rack, inverter, tow hitch, wash, water tank, handicap accessibility, solar, Wi-Fi, or satellite connection, you will be able to take part in numerous activities such as backcountry camping in the Otis Pike High Dune Wilderness, camping at the Fort Wadsworth Campground, tent camping at the Sandy Hook Camp Ground, and family camping at Watch Hill. Experience all of these activities and explore the natural beauty that Brooklyn has to offer!

## Resources

* https://towardsdatascience.com/gpt-3-parameters-and-prompt-design-1a595dc5b405
* https://platform.openai.com/docs/api-reference/completions