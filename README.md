# Marketsim


A Market simulator written in Golang


### Features List

* Implement trading types (http://www.investopedia.com/articles/basics/03/032103.asp)
  * Market Order
  * Limit Order
  * Stop Order
  * All Or None Order
  * Good 'Til Cancelled
  * Day Order


### Architecture 

* Allow the market to be distributed, DNS would be a good starting point to discuss.
* Convert to/from any currency on a trade request
* Accept trades via token in JSON/XML/Whatever
* Data store...? redis?

### Anatomy of a trade
    {
      time: 
      owner:
      type: 
      price:
      quantity:
      commodity:
    }
    
