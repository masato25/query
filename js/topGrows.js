/*
  get diff value of each 2 point and retrun top n
*/
limit = (typeof limit == "undefined"? 3 : limit)
orderby = (typeof orderby == "undefined"? "desc" : orderby)
sortby = (typeof sortby == "undefined"? "Max" : sortby)
t2 = _.map(input, function(res){
  res.Max = 0
  if( res.Values.length == 0){
    return res
  }else{
    storeValues = []
    _.reduce(res.Values, function(lastVal,v){
      value = (isNaN(v.Value)? 0 : v.Value)
      if(lastVal[1] < value){
        tmpVal = (value - lastVal[1]) + lastVal[0]
        lastVal = [tmpVal, value]
      }else{
        storeValues.push(lastVal[0])
        lastVal = [0, value]
      }
      return lastVal
    }, [0,0])
    MaxVals = _.max(storeValues, function(v){
      return v
    })
    res.Max = MaxVals
    return res
  }
})

t3 = _.chain(t2).sortBy(function(res){
  if(orderby == "desc"){
    return - res[sortby]
  }else{
    return res[sortby]
  }

}).first(limit).value()

output = JSON.stringify(t3)
