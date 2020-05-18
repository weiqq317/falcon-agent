using System;
using System.ComponentModel.DataAnnotations;
using System.Net.Http;
using System.Net.Http.Headers;
using Newtonsoft.Json;
namespace dotnet
{
    class Program
    {
     　　public class Metric
        {
            public string endpoint { get; set; }
            public string metric { get; set; }
            public int timestamp { get; set; }
            public int step { get; set; }
            public float value { get; set; }
            public string counterType { get; set; }
            public string tags { get; set; }
            public string addtag { get; set; }

        }

        public static string HttpClientPost(string url,string requestJson)
        {
            try
            {
                string result = string.Empty;
                Uri postUrl = new Uri(url);

                using (HttpContent httpContent = new StringContent(requestJson))
                {
                    httpContent.Headers.ContentType = new MediaTypeHeaderValue("application/json");
                 
                    using (var httpClient = new HttpClient())
                    {
                        httpClient.Timeout = new TimeSpan(0, 0, 60);
                        result = httpClient.PostAsync(url, httpContent).Result.Content.ReadAsStringAsync().Result;
                        
                    }

                }
                return result; 
            }
            catch (Exception e)
            { 
                throw e;
            } 
        }
        public static int ConvertDateTimeInt(System.DateTime time)
        {
            System.DateTime startTime = TimeZone.CurrentTimeZone.ToLocalTime(new System.DateTime(1970, 1, 1));
            return (int)(time - startTime).TotalSeconds;
        }

        static void Main(string[] args)
        {
            //"endpoint": "test-endpoint",
            //"metric": "test-metric",
            //"timestamp": ts,
            //"step": 60,
            //"value": 1,
            //"counterType": "GAUGE",
            //"tags": "location=beijing,service=falcon",
            Metric[] me = new Metric[5];

            for (int i=0; i<5; i++)
            {
                me[i] = new Metric();
                me[i].endpoint = "c# Tester";
                me[i].metric = "vscode.count";
                me[i].timestamp = ConvertDateTimeInt(DateTime.Now);
                me[i].step = 60;
                me[i].value =  0.12345f;
                me[i].counterType = "GAUGE";
                me[i].tags = "vscode=ThinkPadsX230";
                me[i].addtag = "test";

            }



            string json = JsonConvert.SerializeObject(me);
           

           json = HttpClientPost("http://127.0.0.1:1988/v1/push", json);
            Console.WriteLine(json);


            //Console.WriteLine("Hello World!");
           // Console.ReadKey();
        }
    }
}
