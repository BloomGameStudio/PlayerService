using System;
using RestSharp;
namespace HelloWorldApplication {
  class HelloWorld {
    static void Main(string[] args) {
      var client = new RestClient("http://127.0.0.1:1323/player");
      client.Timeout = -1;
      var request = new RestRequest(Method.POST);
      request.AddHeader("Content-Type", "application/json");
      var body = @"{" + "\n" +
      @"    ""UserID"": ""00000000-0000-0000-0000-000000000000""," + "\n" +
      @"    ""Name"": ""User1""," + "\n" +
      @"    ""Layer"": ""layer1""," + "\n" +
      @"    ""Position"": {" + "\n" +
      @"        ""x"": 1," + "\n" +
      @"        ""y"": 2," + "\n" +
      @"        ""z"": 3" + "\n" +
      @"    }," + "\n" +
      @"    ""Rotation"": {" + "\n" +
      @"        ""x"": 4," + "\n" +
      @"        ""y"": 5," + "\n" +
      @"        ""z"": 6" + "\n" +
      @"    }," + "\n" +
      @"    ""Scale"": {" + "\n" +
      @"        ""x"": 7," + "\n" +
      @"        ""y"": 8," + "\n" +
      @"        ""z"": 9" + "\n" +
      @"    }" + "\n" +
      @"}";
      request.AddParameter("application/json", body,  ParameterType.RequestBody);
      IRestResponse response = client.Execute(request);
      Console.WriteLine(response.Content);
    }
  }
}
