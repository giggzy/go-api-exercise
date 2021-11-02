#!/usr/bin/env python
import json
from string import ascii_letters


# generate sample json file for testing
json_file = 'sample.json'

def gen_services():
    records_count = 15
    services = { "services" : []}
    #service_list = services["services"]
    for i in range(records_count):
        """
        "   {
                "name": "NameA",
                "id": "2"
                "desc": "A blah, blah, blah",
                "version_count": 2
                "url": "https://example.com/serviveA"
                "versions": {
                    [
                        {
                            "name": "version1",
                            "id": 1,
                        },
                        {
                            "name": "version2",
                            "id": 2,
                        }
                    ]
                }
            },
        """
        current = ascii_letters[i] 
        name = current + "_Service"
        desc = current + " blah, blah, blah"
        version_count = i + 1
        url = "https://example.com/" + name

        versions = []
        for v in range(version_count):
            versions.append({
                "name": "version_" + str(v),
                "id":   str(v)
            })

        services["services"].append({ 
            "id": str(version_count),
            "name": name,
            "description": desc,
            "url": url,
            "versionCount": version_count,
            "versions": versions
        })
    return services




"""
# print out json file
with open(json_file, 'w') as f:
    for c in ascii_letters:
        service_name += c

        f.write(json.dumps(data, indent=4))
        """


"""
Think about how to generate the json file
Argparse?

type Services struct {
	Services []Service `json:"services"`
}
type Service struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	VersionCount int    `json:"versionCount"`
}
"""
def main():
   services = gen_services()
   with open(json_file, 'w') as f:
       f.write(json.dumps(services, indent=4))
       #f.write(json.dump(json, indent=4))


if __name__ == "__main__":
    main()
