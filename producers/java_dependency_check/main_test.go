package main

import (
	"testing"

	v1 "api/proto/v1"

	"github.com/stretchr/testify/assert"
)

func TestParseIssues(t *testing.T) {
	results := UnmarshalJSON([]byte(exampleOutput))
	issues := parseIssues(results)

	expectedIssue := &v1.Issue{
		Target:      "/src/gen/third_party/go/src/github.com/apache/thrift/lib/netcore/Thrift/Thrift.csproj",
		Type:        "Vulnerable Dependency",
		Title:       "/src/gen/third_party/go/src/github.com/apache/thrift/lib/netcore/Thrift/Thrift.csproj",
		Severity:    v1.Severity_SEVERITY_HIGH,
		Cvss:        7.5,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: ".NET Core 1.0, 1.1, and 2.0 allow an unauthenticated attacker to remotely cause a denial of service attack against a .NET Core web application by improperly parsing certificate data. A denial of service vulnerability exists when .NET Core improperly handles parsing certificate data, aka \".NET CORE Denial Of Service Vulnerability\".",
	}
	assert.Equal(t, issues[0].Target, expectedIssue.Target)
	assert.Equal(t, issues[0].Type, expectedIssue.Type)
	assert.Equal(t, issues[0].Title, expectedIssue.Title)
	assert.Equal(t, issues[0].Severity, expectedIssue.Severity)
	assert.Equal(t, issues[0].Cvss, expectedIssue.Cvss)
	assert.Equal(t, issues[0].Confidence, expectedIssue.Confidence)
	assert.Equal(t, issues[0].Description, expectedIssue.Description)
}

const exampleOutput = `{
	"reportSchema": "1.1",
	"scanInfo": {
		"engineVersion": "5.3.2",
		"dataSource": [{
			"name": "NVD CVE Checked",
			"timestamp": "2020-05-06T12:15:55"
		}, {
			"name": "NVD CVE Modified",
			"timestamp": "2020-05-06T10:01:41"
		}, {
			"name": "VersionCheckOn",
			"timestamp": "2020-05-06T12:15:55"
		}]
	},
	"projectInfo": {
		"name": "testproj",
		"reportDate": "2020-05-06T12:16:00.051917Z",
		"credits": {
			"NVD": "This report contains data retrieved from the National Vulnerability Database: http://nvd.nist.gov",
			"NPM": "This report may contain data retrieved from the NPM Public Advisories: https://www.npmjs.com/advisories",
			"RETIREJS": "This report may contain data retrieved from the RetireJS community: https://retirejs.github.io/retire.js/",
			"OSSINDEX": "This report may contain data retrieved from the Sonatype OSS Index: https://ossindex.sonatype.org"
		}
	},
	"dependencies": [{
		"isVirtual": false,
		"fileName": "Client.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/tutorial\/netcore\/Client\/Client.csproj",
		"md5": "a69a91e7f3b7f983f4368b5680b8e1e4",
		"sha1": "7cd4dfa9632b96f125cc6fa3bfc4990dab08edf0",
		"sha256": "1c9346a726417cafae77d170a977c2d2110f11ab2885da0823d968694c7cf499",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Client"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Client"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "CsharpClient.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/tutorial\/csharp\/CsharpClient\/CsharpClient.csproj",
		"md5": "09ba517ffea5997ab93ccc04299d7d1a",
		"sha1": "17fc7bd3b4a4576329c0225599d861350277dc84",
		"sha256": "ff7bbcd9af07cadf7b4f1fb3757acbe0bc0f36144fdead6ff39129cafdbe9153",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "CsharpClient"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "CsharpClient"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "CsharpServer.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/tutorial\/csharp\/CsharpServer\/CsharpServer.csproj",
		"md5": "3a409a1cbbe8133aaa433e13cd220b8b",
		"sha1": "99c85c946d68ade035ddecfc1ab1e6f4dff51211",
		"sha256": "d963d488dfb6ea826a352c557ed21ca3bc68cd6a2c197ba03d3e55cc3d05294a",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "CsharpServer"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "CsharpServer"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "Interfaces.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/tutorial\/netcore\/Interfaces\/Interfaces.csproj",
		"md5": "a725657af9a89944854039157072764a",
		"sha1": "a18c1e71a7ae32335b173c8513fe8915be4e7fc6",
		"sha256": "b550b4d39a80aa481273edc3257a9d7c032d690370517ccdb49f449c2833af23",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Interfaces"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Interfaces"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "JSONTest.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/csharp\/test\/JSON\/JSONTest.csproj",
		"md5": "4016b0ea988d95529d760f58bec143f2",
		"sha1": "8c2a718ff037dc4a30deb9891d19ac4df049aa2d",
		"sha256": "609cdcab43401b296e77c52c1e22798c0712cebdc1628a3c6f2e736a1e3667ff",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "JSONTest"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "JSONTest"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": true,
		"fileName": "Microsoft.AspNetCore.Server.IISIntegration:[2.0,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/tutorial\/netcore\/Server\/Server.csproj",
		"relatedDependencies": [{
			"isVirtual": true,
			"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/tutorial\/netcore\/Server\/Server.csproj",
			"packageIds": [{
				"id": "pkg:nuget/Microsoft.AspNetCore.Server.Kestrel@%5B2.0%2C%29"
			}]
		}],
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "AspNetCore.Server.IISIntegration"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft.AspNetCore.Server.IISIntegration"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "AspNetCore"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[2.0,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/Microsoft.AspNetCore.Server.IISIntegration@%5B2.0%2C%29",
			"confidence": "HIGHEST"
		}],
		"vulnerabilityIds": [{
			"id": "cpe:2.3:a:microsoft:project_server:2.0:*:*:*:*:*:*:*",
			"confidence": "LOW"
		}, {
			"id": "cpe:2.3:a:microsoft:server:2.0:*:*:*:*:*:*:*",
			"confidence": "LOW"
		}]
	}, {
		"isVirtual": true,
		"fileName": "Microsoft.AspNetCore:[2.0,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Thrift\/Thrift.csproj",
		"relatedDependencies": [{
			"isVirtual": true,
			"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/tutorial\/netcore\/Server\/Server.csproj",
			"packageIds": [{
				"id": "pkg:nuget/Microsoft.AspNetCore@%5B2.0%2C%29"
			}]
		}],
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft.AspNetCore"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "AspNetCore"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[2.0,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/Microsoft.AspNetCore@%5B2.0%2C%29",
			"confidence": "HIGHEST"
		}],
		"vulnerabilityIds": [{
			"id": "cpe:2.3:a:microsoft:aspnetcore:2.0:*:*:*:*:*:*:*",
			"confidence": "HIGHEST",
			"url": "https:\/\/nvd.nist.gov\/vuln\/search\/results?form_type=Advanced&results_type=overview&search_type=all&cpe_vendor=cpe%3A%2F%3Amicrosoft&cpe_product=cpe%3A%2F%3Amicrosoft%3Aaspnetcore&cpe_version=cpe%3A%2F%3Amicrosoft%3Aaspnetcore%3A2.0"
		}],
		"vulnerabilities": [{
			"source": "NVD",
			"name": "CVE-2017-11770",
			"severity": "HIGH",
			"cvssv2": {
				"score": 5.0,
				"accessVector": "NETWORK",
				"accessComplexity": "LOW",
				"authenticationr": "NONE",
				"confidentialImpact": "NONE",
				"integrityImpact": "NONE",
				"availabilityImpact": "PARTIAL",
				"severity": "MEDIUM"
			},
			"cvssv3": {
				"baseScore": 7.5,
				"attackVector": "NETWORK",
				"attackComplexity": "LOW",
				"privilegesRequired": "NONE",
				"userInteraction": "NONE",
				"scope": "UNCHANGED",
				"confidentialityImpact": "NONE",
				"integrityImpact": "NONE",
				"availabilityImpact": "HIGH",
				"baseSeverity": "HIGH"
			},
			"cwes": ["CWE-295"],
			"description": ".NET Core 1.0, 1.1, and 2.0 allow an unauthenticated attacker to remotely cause a denial of service attack against a .NET Core web application by improperly parsing certificate data. A denial of service vulnerability exists when .NET Core improperly handles parsing certificate data, aka \".NET CORE Denial Of Service Vulnerability\".",
			"notes": "",
			"references": [{
				"source": "CONFIRM",
				"url": "https:\/\/portal.msrc.microsoft.com\/en-US\/security-guidance\/advisory\/CVE-2017-11770",
				"name": "https:\/\/portal.msrc.microsoft.com\/en-US\/security-guidance\/advisory\/CVE-2017-11770"
			}, {
				"source": "BID",
				"url": "http:\/\/www.securityfocus.com\/bid\/101710",
				"name": "101710"
			}, {
				"source": "REDHAT",
				"url": "https:\/\/access.redhat.com\/errata\/RHSA-2017:3248",
				"name": "RHSA-2017:3248"
			}, {
				"source": "SECTRACK",
				"url": "http:\/\/www.securitytracker.com\/id\/1039787",
				"name": "1039787"
			}],
			"vulnerableSoftware": [{
				"software": {
					"id": "cpe:2.3:a:microsoft:aspnetcore:1.0:*:*:*:*:*:*:*"
				}
			}, {
				"software": {
					"id": "cpe:2.3:a:microsoft:aspnetcore:2.0:*:*:*:*:*:*:*",
					"vulnerabilityIdMatched": "true"
				}
			}, {
				"software": {
					"id": "cpe:2.3:a:microsoft:aspnetcore:1.1:*:*:*:*:*:*:*"
				}
			}]
		}, {
			"source": "NVD",
			"name": "CVE-2017-11883",
			"severity": "HIGH",
			"cvssv2": {
				"score": 5.0,
				"accessVector": "NETWORK",
				"accessComplexity": "LOW",
				"authenticationr": "NONE",
				"confidentialImpact": "NONE",
				"integrityImpact": "NONE",
				"availabilityImpact": "PARTIAL",
				"severity": "MEDIUM"
			},
			"cvssv3": {
				"baseScore": 7.5,
				"attackVector": "NETWORK",
				"attackComplexity": "LOW",
				"privilegesRequired": "NONE",
				"userInteraction": "NONE",
				"scope": "UNCHANGED",
				"confidentialityImpact": "NONE",
				"integrityImpact": "NONE",
				"availabilityImpact": "HIGH",
				"baseSeverity": "HIGH"
			},
			"cwes": ["NVD-CWE-noinfo"],
			"description": ".NET Core 1.0, 1.1, and 2.0 allow an unauthenticated attacker to remotely cause a denial of service attack against a .NET Core web application by improperly handling web requests, aka \".NET CORE Denial Of Service Vulnerability\".",
			"notes": "",
			"references": [{
				"source": "BID",
				"url": "http:\/\/www.securityfocus.com\/bid\/101835",
				"name": "101835"
			}, {
				"source": "SECTRACK",
				"url": "http:\/\/www.securitytracker.com\/id\/1039793",
				"name": "1039793"
			}, {
				"source": "CONFIRM",
				"url": "https:\/\/portal.msrc.microsoft.com\/en-US\/security-guidance\/advisory\/CVE-2017-11883",
				"name": "https:\/\/portal.msrc.microsoft.com\/en-US\/security-guidance\/advisory\/CVE-2017-11883"
			}],
			"vulnerableSoftware": [{
				"software": {
					"id": "cpe:2.3:a:microsoft:aspnetcore:1.0:*:*:*:*:*:*:*"
				}
			}, {
				"software": {
					"id": "cpe:2.3:a:microsoft:aspnetcore:2.0:*:*:*:*:*:*:*",
					"vulnerabilityIdMatched": "true"
				}
			}, {
				"software": {
					"id": "cpe:2.3:a:microsoft:aspnetcore:1.1:*:*:*:*:*:*:*"
				}
			}]
		}]
	}, {
		"isVirtual": true,
		"fileName": "Microsoft.Extensions.Configuration.FileExtensions:[2.0,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/tutorial\/netcore\/Server\/Server.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Extensions.Configuration.FileExtensions"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft.Extensions.Configuration.FileExtensions"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Extensions"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[2.0,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/Microsoft.Extensions.Configuration.FileExtensions@%5B2.0%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "Microsoft.Extensions.Logging.Console:[2.0,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Thrift\/Thrift.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft.Extensions.Logging.Console"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Extensions"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Extensions.Logging.Console"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[2.0,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/Microsoft.Extensions.Logging.Console@%5B2.0%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "Microsoft.Extensions.Logging.Debug:[2.0,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Thrift\/Thrift.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Extensions.Logging.Debug"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft.Extensions.Logging.Debug"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Extensions"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[2.0,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/Microsoft.Extensions.Logging.Debug@%5B2.0%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "Microsoft.Extensions.Logging:[2.0,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Thrift\/Thrift.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "Microsoft.Extensions.Logging"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Extensions"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Extensions.Logging"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[2.0,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/Microsoft.Extensions.Logging@%5B2.0%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": false,
		"fileName": "MultiplexClient.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/csharp\/test\/Multiplex\/Client\/MultiplexClient.csproj",
		"md5": "d070c60dee2eec0b9daada4027407064",
		"sha1": "a2edefe53593ff6756db807259d187f677ed6760",
		"sha256": "f3f5ed3d6b9c06beb6d1206c01c2d7a127ca328a3cc5293b900d3018ca893b9a",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "MultiplexClient"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "MultiplexClient"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "MultiplexServer.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/csharp\/test\/Multiplex\/Server\/MultiplexServer.csproj",
		"md5": "ded7e48e6c4ee5a04d9e163cf951c117",
		"sha1": "c3c5ff6ce45ca413aee8fa0164f585e2471aec32",
		"sha256": "8f8ca4b74a35bc9e6c79063ef2a9c77561129a33076504119b809d7944fbe288",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "MultiplexServer"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "MultiplexServer"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "RebusSample.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/contrib\/Rebus\/RebusSample.csproj",
		"md5": "06abcaa9eb5c09c06c6d9cf1f6f3e782",
		"sha1": "4634909b317398e6a6759476a01bdd4cdbdc725f",
		"sha256": "5d1adb8c5fbc67d8a0bdb68d2f544039e9b612046cc5e01e7555e86d57f5ca6c",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "RebusSample"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "RebusSample"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "Server.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/tutorial\/netcore\/Server\/Server.csproj",
		"md5": "515b56ec9f5853b3f5ed16c662137dab",
		"sha1": "5c73a7435bb7df998c45104599f18ee34f4f560f",
		"sha256": "0464e02d6a79d7427550527ff915805e7c662ed41d2b7360bf75ad7596d999d1",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Server"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Server"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": true,
		"fileName": "System.IO.Pipes:[4.3,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Thrift\/Thrift.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "System"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "IO.Pipes"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "System.IO.Pipes"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "IO"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[4.3,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/System.IO.Pipes@%5B4.3%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "System.Net.Http.WinHttpHandler:4.4.0",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Thrift\/Thrift.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "System"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Net"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Net.Http.WinHttpHandler"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "System.Net.Http.WinHttpHandler"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "4.4.0"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/System.Net.Http.WinHttpHandler@4.4.0",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "System.Net.Http.WinHttpHandler:[4.4,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/test\/netcore\/ThriftTest\/ThriftTest.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "System"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Net"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Net.Http.WinHttpHandler"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "System.Net.Http.WinHttpHandler"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[4.4,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/System.Net.Http.WinHttpHandler@%5B4.4%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "System.Net.NameResolution:[4.3,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Thrift\/Thrift.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "System"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Net"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Net.NameResolution"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "System.Net.NameResolution"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[4.3,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/System.Net.NameResolution@%5B4.3%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "System.Net.Requests:[4.3,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Thrift\/Thrift.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "System"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Net"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "System.Net.Requests"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Net.Requests"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[4.3,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/System.Net.Requests@%5B4.3%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "System.Net.Security:[4.3,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Thrift\/Thrift.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "System"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Net"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Net.Security"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "System.Net.Security"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[4.3,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/System.Net.Security@%5B4.3%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "System.Runtime.Serialization.Primitives:[4.3,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/test\/netcore\/ThriftTest\/ThriftTest.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "System"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "System.Runtime.Serialization.Primitives"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Runtime"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Runtime.Serialization.Primitives"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[4.3,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/System.Runtime.Serialization.Primitives@%5B4.3%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "System.ServiceModel.Primitives:[4.1.0,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Tests\/Thrift.PublicInterfaces.Compile.Tests\/Thrift.PublicInterfaces.Compile.Tests.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "System"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "ServiceModel.Primitives"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "System.ServiceModel.Primitives"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "ServiceModel"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[4.1.0,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/System.ServiceModel.Primitives@%5B4.1.0%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "System.ServiceModel.Primitives:[4.4,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/test\/netcore\/ThriftTest\/ThriftTest.csproj",
		"relatedDependencies": [{
			"isVirtual": true,
			"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/tutorial\/netcore\/Interfaces\/Interfaces.csproj",
			"packageIds": [{
				"id": "pkg:nuget/System.ServiceModel.Primitives@%5B4.4%2C%29"
			}]
		}],
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "System"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "ServiceModel.Primitives"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "System.ServiceModel.Primitives"
			}, {
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "ServiceModel"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[4.4,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/System.ServiceModel.Primitives@%5B4.4%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": true,
		"fileName": "System.Threading:[4.3,)",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/test\/netcore\/ThriftTest\/ThriftTest.csproj",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "System"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "MEDIUM",
				"source": "msbuild",
				"name": "id",
				"value": "Threading"
			}, {
				"type": "product",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "id",
				"value": "System.Threading"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "HIGHEST",
				"source": "msbuild",
				"name": "version",
				"value": "[4.3,)"
			}]
		},
		"packages": [{
			"id": "pkg:nuget\/System.Threading@%5B4.3%2C%29",
			"confidence": "HIGHEST"
		}]
	}, {
		"isVirtual": false,
		"fileName": "Thrift.45.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/csharp\/src\/Thrift.45.csproj",
		"md5": "89ef4542b5b2967ad2f913cb6dc8b615",
		"sha1": "6d8082655112e103881f7ad541c5692774c08b1c",
		"sha256": "a632d1dc90e167fa9a4e9682ffef5c846956aa4f65cd85871af0771863d1f9fa",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Thrift.45"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Thrift.45"
			}],
			"versionEvidence": [{
				"type": "version",
				"confidence": "MEDIUM",
				"source": "file",
				"name": "version",
				"value": "45"
			}, {
				"type": "version",
				"confidence": "MEDIUM",
				"source": "file",
				"name": "name",
				"value": "Thrift.45"
			}]
		}
	}, {
		"isVirtual": false,
		"fileName": "Thrift.PublicInterfaces.Compile.Tests.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Tests\/Thrift.PublicInterfaces.Compile.Tests\/Thrift.PublicInterfaces.Compile.Tests.csproj",
		"md5": "3033c0975524750b9f6caf07f8b25944",
		"sha1": "44be129b1d31d29861239d45fe46951faae1fd1f",
		"sha256": "202c293ddc0e738736e23421779ceb33a09f957fa6f5982aa8eb467251d88ce7",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Thrift.PublicInterfaces.Compile.Tests"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Thrift.PublicInterfaces.Compile.Tests"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "Thrift.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/csharp\/src\/Thrift.csproj",
		"md5": "f966536175601ec180e9ed311bc83efc",
		"sha1": "41349987388ec649d29dbae1996dc4cb91fefdd0",
		"sha256": "a100497623b01f6b0697b3edb0523c533e851845668cde70d19604356f1901cb",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Thrift"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Thrift"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "Thrift.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/netcore\/Thrift\/Thrift.csproj",
		"md5": "f2046ee056185cbb37a75ea727e06f09",
		"sha1": "7576338264596831b90e3e7d264e06bd9c34a37c",
		"sha256": "851b25c85665047c5f3e061ead145657259523cf027fbb94163ab7f717d0e2dc",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Thrift"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "Thrift"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "ThriftMSBuildTask.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/csharp\/ThriftMSBuildTask\/ThriftMSBuildTask.csproj",
		"md5": "90390df6655f9246c0f62640c3258b1d",
		"sha1": "25e18db2eeb035870c8dfe2682c837124fa824cf",
		"sha256": "47a9a378154bfebac48876767f63d10dd7bd0872a5180e44dbefab9e37793122",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "ThriftMSBuildTask"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "ThriftMSBuildTask"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "ThriftMVCTest.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/lib\/csharp\/test\/ThriftMVCTest\/ThriftMVCTest.csproj",
		"md5": "4126527d200789b0d00b201a29df518c",
		"sha1": "71e6236ee0f444cf865d94539e0e7393bb22c833",
		"sha256": "d0e01b105de27f3ecc010323461cedb432e5d22e5afa22d76747a87e27b79bf0",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "ThriftMVCTest"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "ThriftMVCTest"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "ThriftTest.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/test\/netcore\/ThriftTest\/ThriftTest.csproj",
		"md5": "4e901f1ffaaef35c008cb9488ce94922",
		"sha1": "dc64fa9cf953c09312fc65c7a32ca3cc571726d6",
		"sha256": "04a18beaf2ebfd92ec3b7862e21da6a392939943a3a8e483b15aa810df2ce564",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "ThriftTest"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "ThriftTest"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "ThriftTest.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/test\/csharp\/ThriftTest.csproj",
		"md5": "0e9352ae9ccde6793969526506205dbb",
		"sha1": "010ff1fafe07d9b7b22f0d027e7fe1b07f35af30",
		"sha256": "da335814a51cee513dc66bd162578ac17aefa949ddd3b23c727aea55bc6613bd",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "ThriftTest"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "ThriftTest"
			}],
			"versionEvidence": []
		}
	}, {
		"isVirtual": false,
		"fileName": "ThriftZMQ.csproj",
		"filePath": "\/src\/gen\/third_party\/go\/src\/github.com\/apache\/thrift\/contrib\/zeromq\/csharp\/ThriftZMQ.csproj",
		"md5": "80b6cd5d8bfb9b0170f4ff46006e65f2",
		"sha1": "e0ab0c62cca1af265d20cd0371a21a4aab7efe9e",
		"sha256": "784aff660f55b74af0f7329f174a81f552ee2de2b08d34d74af63af7086e51ca",
		"evidenceCollected": {
			"vendorEvidence": [{
				"type": "vendor",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "ThriftZMQ"
			}],
			"productEvidence": [{
				"type": "product",
				"confidence": "HIGH",
				"source": "file",
				"name": "name",
				"value": "ThriftZMQ"
			}],
			"versionEvidence": []
		}
	}]
}`
