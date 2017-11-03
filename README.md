# webHooksAndCurrencies
A failed tale of webhooks, currencies, failed heroku deployments, a love (wait, no, hate) triangle between me, govendor and godep, and of apps not being noticed in buildpacks

Me and some of the others in the class have sat for actual days now trying to deploy *something* to heroku to no avail.
The plan was to get it to heroku and then keep making changes to make things work. However, even with the help of other
classmates who have gotten this working, and trying it on all of our computers, this doesn't work. Thus, what we have is
frankly very little.

We've got an API for webhooks. It works fairly well. Running userWebHookAPI.go will prompt you to type in the
URL for a webhook. Then you're asked to add in the data for it. After this, we go through the function invoking all
of the webhooks currently saved in the database, and then pushing out that info to each webhook.

We also managed to connect to a mongodb database, and it works well. We can put in and collect info both for webhooks
and for the currency rates. Though we don't have a timer for that at the moment, every time you run it it collects
an instance of all the rates and saves it in the database.

The average of three days function doesn't work. The remains of it can be found in currencyTicker_db.go on line 200.
Basically, the issue came with extracting the float for the specific target currency.

Everything got put on the back burner for trying to get our god damned project onto heroku. I just included everything into
the git repo, so you'll be able to see that we've done something at least.

I've spent several days now with far too little sleep, neglecting other group projects I have. I give up. Go wins this round.



Go makes a lot of claims, people have said it's beautiful, and their twitter even
says they promise to reignite your love for programming, but frankly I just find it
extremely ugly and impractical, as well as hard to look through, which I suppose
defeats the entire point of the language in the first place
