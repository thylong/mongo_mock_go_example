Mongo Mock example in Go
========================

.. image:: https://travis-ci.org/thylong/mongo_mock_go_example.svg?branch=master
    :target: https://travis-ci.org/thylong/mongo_mock_go_example

This repository contains a simple way to mock Mgo driver.

If you're interested in it, you can find an article here : http://thylong.com/golang/2016/mocking-mongo-in-golang/


Installing
==========

.. code-block:: console

    $ git clone git@github.com:thylong/mongo_mock_go_example.git

Quick-start
===========

If you want to test only the mock, you can run it this way :

.. code-block:: console

    $ go get github.com/thylong/mongo_mock_go_example
    $ go run main.go

If you want to test both the connection to Mongo and the mock, you can use the docker-compose file.


Nota Bene
=========

The code was tested with versions 1.7.* of Go.
