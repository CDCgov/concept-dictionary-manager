package main

var CodeSystemsMapping = string(`{
    "settings": {
        "index": {
            "creation_date": "1488923647218",
            "number_of_shards": "5",
            "number_of_replicas": "1",
            "uuid": "2eR2Kvt4TvqWSkwyaYccYA",
            "version": {
                "created": "5020299"
            },
            "provided_name": "code_systems"
        }
    },
    "mappings": {
        "code_system": {
            "properties": {
                "acquiredDate": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "assigningAuthorityId": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "assigningAuthorityReleaseDate": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "assigningAuthorityVersionName": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "codeSystemCode": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "definitionstring": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "distributionSourceId": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "distributionSourceReleaseDate": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "distributionSourceVersionName": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "effectiveDate": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "expiryDate": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "hl70396Identifier": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "id": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "lastRevisionDate": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "legacyFlag": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "oid": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "sdoCreateDate": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "sourceUrl": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "status": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "statusDate": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "version": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                },
                "versionDescription": {
                    "type": "string",
                    "fields": {
                        "keyword": {
                            "type": "keyword",
                            "ignore_above": 256
                        }
                    }
                }
            }
        }
    }
}`)

var CodeSystemJSON = string(`{
        "oid": "2.16.840.1.113883.5.42",
        "id": "87BE6F37-6967-DE11-9B52-0015173D1785",
        "name": "EntityHandling",
        "definitionText": "Entity Handling (HL7) - List of codes to indicate any special handling required for the Non-Person Living Subject (NPLS).",
        "status": "Published",
        "statusDate": "Jul 2, 2009 8:34:07 PM",
        "version": "20090501",
        "versionDescription": "hl7_rimrepos-2.26.2",
        "acquiredDate": "May 1, 2009 12:00:00 AM",
        "assigningAuthorityVersionName": "hl7_rimrepos-2.26.2",
        "assigningAuthorityReleaseDate": "May 1, 2009 12:00:00 AM",
        "distributionSourceVersionName": "May 2009 Ballot",
        "distributionSourceReleaseDate": "May 1, 2009 12:00:00 AM",
        "distributionSourceId": "8A1A50DA-0CAC-DD11-BB1F-00188B398520",
        "lastRevisionDate": "May 1, 2009 12:00:00 AM",
        "assigningAuthorityId": "31F87B81-0CAC-DD11-BB1F-00188B398520",
        "codeSystemCode": "PH_EntityHandling_HL7_V3",
        "sourceUrl": "http://hl7projects.hl7.nscee.edu/frs/download.php/622/hl7_rimrepos-2.26.2.zip",
        "hl70396Identifier": "ENTITYHDLG",
        "legacyFlag": false
}`)
